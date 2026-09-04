package apps

import (
	"context"
	"os"

	"github.com/MamangRust/monolith-graphql-pointofsale-order/handler"
	mencache "github.com/MamangRust/monolith-graphql-pointofsale-order/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-order/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-order/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"
	pbroot "github.com/MamangRust/monolith-graphql-pointofsale-pb"
	pbcashier "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pbproduct "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"

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
	productAddr := getEnv("GRPC_PRODUCT_ADDR", "localhost:50059")
	orderItemAddr := getEnv("GRPC_ORDERITEM_ADDR", "localhost:50057")

	srv.Logger.Info("Connecting to gRPC microservices from Order microservice",
		zap.String("cashier", cashierAddr),
		zap.String("merchant", merchantAddr),
		zap.String("product", productAddr),
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

	productConn, err := server.NewGRPCClient(productAddr)
	if err != nil {
		cashierConn.Close()
		merchantConn.Close()
		return nil, err
	}

	orderItemConn, err := server.NewGRPCClient(orderItemAddr)
	if err != nil {
		cashierConn.Close()
		merchantConn.Close()
		productConn.Close()
		return nil, err
	}

	go func() {
		<-srv.Ctx.Done()
		srv.Logger.Info("Closing gRPC client connections in Order microservice")
		cashierConn.Close()
		merchantConn.Close()
		productConn.Close()
		orderItemConn.Close()
	}()

	cashierClient := pbcashier.NewCashierQueryServiceClient(cashierConn)
	merchantClient := pbmerchant.NewMerchantQueryServiceClient(merchantConn)
	productClient := pbproduct.NewProductQueryServiceClient(productConn)
	orderItemClient := pbroot.NewOrderItemServiceClient(orderItemConn)

	repos := repository.NewRepositories(srv.DB, cashierClient, merchantClient, productClient, orderItemClient)
	traceLoggerObservability := observability.NewTraceLoggerObservability(srv.Logger)

	cacheMetrics, err := observability.NewCacheMetrics("order")
	if err != nil {
		return nil, err
	}

	cacheStore := cache.NewCacheStore(srv.Redis, srv.Logger, cacheMetrics)
	mencacheObj := mencache.NewMencache(cacheStore)

	services := service.NewService(&service.Deps{
		Mencache:      mencacheObj,
		Ctx:           context.Background(),
		Repositories:  repos,
		Logger:        srv.Logger,
		Observability: traceLoggerObservability,
	})

	handlers := handler.NewHandler(&handler.Deps{
		Service: services,
		Logger:  srv.Logger,
	})

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterOrderQueryServiceServer(gs, handlers.Order)
		pb.RegisterOrderCommandServiceServer(gs, handlers.OrderCommand)
		pb.RegisterOrderStatsServiceServer(gs, handlers.OrderStats)
	}

	return srv, nil
}
