package apps

import (
	"os"

	mencache "github.com/MamangRust/monolith-graphql-pointofsale-auth/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-auth/handler"
	"github.com/MamangRust/monolith-graphql-pointofsale-auth/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-auth/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/auth"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/hash"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/kafka"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/outbox"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	tokenManager, err := auth.NewManager(viper.GetString("SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepositories(srv.DB)
	hash := hash.NewHashingPassword()
	// Kafka bersifat opsional: tanpa KAFKA_BROKERS (mis. E2E lokal tanpa kafka)
	// service tetap jalan dan event email di-skip (guard s.kafka != nil).
	var myKafka *kafka.Kafka
	if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		myKafka = kafka.NewKafka(srv.Logger, []string{brokers})
	}

	mencacheObj := mencache.NewMencache(srv.CacheStore)

	// Phase 6 — transactional outbox: producers enqueue email events inside the
	// business transaction; the relay publishes them durably with retry/DLQ.
	outboxService := outbox.NewOutboxService(srv.DB, myKafka, srv.Logger)

	services := service.NewService(&service.Deps{
		Mencache:      mencacheObj,
		Repositories:  repos,
		Hash:          hash,
		Token:         tokenManager,
		Logger:        srv.Logger,
		Kafka:         myKafka,
		Pool:          srv.DBPool,
		Outbox:        outboxService,
		Observability: observability.NewTraceLoggerObservability(srv.Logger),
	})

	handlers := handler.NewHandler(&handler.Deps{
		Service: services,
		Logger:  srv.Logger,
	})

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterAuthServiceServer(gs, handlers.Auth)
	}

	// Start the outbox relay so events committed with the business writes are
	// published to Kafka with durable retry and dead-letter semantics.
	go outboxService.Start(srv.Ctx, outbox.OutboxRelayInterval, outbox.OutboxRelayBatchSize)

	return srv, nil
}
