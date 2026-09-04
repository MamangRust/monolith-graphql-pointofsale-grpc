package apps

import (
	"context"
	"os"

	mencache "github.com/MamangRust/monolith-graphql-pointofsale-merchant/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-merchant/handler"
	"github.com/MamangRust/monolith-graphql-pointofsale-merchant/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-merchant/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/kafka"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/outbox"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pbmerchant_document "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant_document"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	pbuser "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"

)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	userAddr := os.Getenv("GRPC_USER_ADDR")
	if userAddr == "" {
		userAddr = "localhost:50053"
	}

	srv.Logger.Info("Connecting to User service via gRPC", zap.String("addr", userAddr))
	userConn, err := server.NewGRPCClient(userAddr)
	if err != nil {
		return nil, err
	}

	go func() {
		<-srv.Ctx.Done()
		srv.Logger.Info("Closing merchant service remote gRPC connections")
		userConn.Close()
	}()

	userClient := pbuser.NewUserQueryServiceClient(userConn)

	repos := repository.NewRepositories(srv.DB, userClient)
	// Kafka bersifat opsional: tanpa KAFKA_BROKERS (mis. E2E lokal tanpa kafka)
	// service tetap jalan dan event email di-skip (guard s.kafka != nil).
	var myKafka *kafka.Kafka
	if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		myKafka = kafka.NewKafka(srv.Logger, []string{brokers})
	}
	traceLoggerObservability := observability.NewTraceLoggerObservability(srv.Logger)

	mencacheObj := mencache.NewMencache(srv.CacheStore)

	// Phase 6 — transactional outbox: producers enqueue email events inside the
	// business transaction; the relay publishes them durably with retry/DLQ.
	outboxService := outbox.NewOutboxService(srv.DB, myKafka, srv.Logger)

	services := service.NewService(&service.Deps{
		Ctx:           context.Background(),
		Mencache:      mencacheObj,
		Kafka:         myKafka,
		Repositories:  repos,
		Pool:          srv.DBPool,
		Outbox:        outboxService,
		Logger:        srv.Logger,
		Observability: traceLoggerObservability,
	})

	handlers := handler.NewHandler(&handler.Deps{
		Service: services,
		Logger:  srv.Logger,
	})

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterMerchantQueryServiceServer(gs, handlers.Merchant)
		pb.RegisterMerchantCommandServiceServer(gs, handlers.MerchantCommand)
		pbmerchant_document.RegisterMerchantDocumentQueryServiceServer(gs, handlers.MerchantDocument)
		pbmerchant_document.RegisterMerchantDocumentCommandServiceServer(gs, handlers.MerchantDocumentCommand)
	}

	// Start the outbox relay so events committed with the business writes are
	// published to Kafka with durable retry and dead-letter semantics.
	go outboxService.Start(srv.Ctx, outbox.OutboxRelayInterval, outbox.OutboxRelayBatchSize)

	return srv, nil
}
