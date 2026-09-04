package apps

import (
	"os"

	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/kafka"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/outbox"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	mencache "github.com/MamangRust/monolith-graphql-pointofsale-transacton/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-transacton/handler"
	"github.com/MamangRust/monolith-graphql-pointofsale-transacton/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-transacton/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/transaction"
	pbroot "github.com/MamangRust/monolith-graphql-pointofsale-pb"
	pbcashier "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pborder "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"

)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	cashierAddr := getEnv("GRPC_CASHIER_ADDR", "localhost:50055")
	merchantAddr := getEnv("GRPC_MERCHANT_ADDR", "localhost:50056")
	orderAddr := getEnv("GRPC_ORDER_ADDR", "localhost:50058")
	orderItemAddr := getEnv("GRPC_ORDERITEM_ADDR", "localhost:50057")

	srv.Logger.Info("Connecting to gRPC microservices from Transaction microservice",
		zap.String("cashier", cashierAddr),
		zap.String("merchant", merchantAddr),
		zap.String("order", orderAddr),
		zap.String("order_item", orderItemAddr),
	)

	cashierConn, err := server.NewGRPCClient(cashierAddr)
	if err != nil {
		return nil, err
	}

	merchantConn, err := server.NewGRPCClient(merchantAddr)
	if err != nil {
		cashierConn.Close()
		return nil, err
	}

	orderConn, err := server.NewGRPCClient(orderAddr)
	if err != nil {
		cashierConn.Close()
		merchantConn.Close()
		return nil, err
	}

	orderItemConn, err := server.NewGRPCClient(orderItemAddr)
	if err != nil {
		cashierConn.Close()
		merchantConn.Close()
		orderConn.Close()
		return nil, err
	}

	go func() {
		<-srv.Ctx.Done()
		srv.Logger.Info("Closing gRPC client connections in Transaction microservice")
		cashierConn.Close()
		merchantConn.Close()
		orderConn.Close()
		orderItemConn.Close()
	}()

	cashierClient := pbcashier.NewCashierQueryServiceClient(cashierConn)
	merchantClient := pbmerchant.NewMerchantQueryServiceClient(merchantConn)
	orderClient := pborder.NewOrderQueryServiceClient(orderConn)
	orderItemClient := pbroot.NewOrderItemServiceClient(orderItemConn)

	repos := repository.NewRepositories(srv.DB, cashierClient, merchantClient, orderClient, orderItemClient)
	// Kafka bersifat opsional: tanpa KAFKA_BROKERS (mis. E2E lokal tanpa kafka)
	// service tetap jalan dan event email di-skip (guard s.kafka != nil).
	var myKafka *kafka.Kafka
	if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		myKafka = kafka.NewKafka(srv.Logger, []string{brokers})
	}

	mencacheObj := mencache.NewMencache(srv.CacheStore)

	traceLoggerObservability := observability.NewTraceLoggerObservability(srv.Logger)

	// Phase 6 — transactional outbox: producers enqueue email events inside the
	// business transaction; the relay publishes them durably with retry/DLQ.
	outboxService := outbox.NewOutboxService(srv.DB, myKafka, srv.Logger)

	services := service.NewService(&service.Deps{
		Mencache:      mencacheObj,
		Repositories:  repos,
		Logger:        srv.Logger,
		Kafka:         myKafka,
		Pool:          srv.DBPool,
		Outbox:        outboxService,
		Observability: traceLoggerObservability,
	})

	handlers := handler.NewHandler(&handler.Deps{
		Service: services,
		Logger:  srv.Logger,
	})

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterTransactionQueryServiceServer(gs, handlers.Transaction)
		pb.RegisterTransactionCommandServiceServer(gs, handlers.TransactionCommand)
		pb.RegisterTransactionStatsServiceServer(gs, handlers.TransactionStats)
	}

	// Start the outbox relay so events committed with the business writes are
	// published to Kafka with durable retry and dead-letter semantics.
	go outboxService.Start(srv.Ctx, outbox.OutboxRelayInterval, outbox.OutboxRelayBatchSize)

	return srv, nil
}
