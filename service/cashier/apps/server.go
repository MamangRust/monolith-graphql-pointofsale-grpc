package apps

import (
	"context"
	"os"

	"github.com/MamangRust/monolith-graphql-pointofsale-cashier/handler"
	mencache "github.com/MamangRust/monolith-graphql-pointofsale-cashier/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-cashier/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-cashier/service"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	pbuser "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"

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
	merchantAddr := os.Getenv("GRPC_MERCHANT_ADDR")
	if merchantAddr == "" {
		merchantAddr = "localhost:50056"
	}

	srv.Logger.Info("Connecting to User service via gRPC", zap.String("addr", userAddr))
	userConn, err := server.NewGRPCClient(userAddr)
	if err != nil {
		return nil, err
	}

	srv.Logger.Info("Connecting to Merchant service via gRPC", zap.String("addr", merchantAddr))
	merchantConn, err := server.NewGRPCClient(merchantAddr)
	if err != nil {
		userConn.Close()
		return nil, err
	}

	go func() {
		<-srv.Ctx.Done()
		srv.Logger.Info("Closing cashier service remote gRPC connections")
		userConn.Close()
		merchantConn.Close()
	}()

	userClient := pbuser.NewUserQueryServiceClient(userConn)
	merchantClient := pbmerchant.NewMerchantQueryServiceClient(merchantConn)

	repos := repository.NewRepositories(srv.DB, userClient, merchantClient)
	traceLoggerObservability := observability.NewTraceLoggerObservability(srv.Logger)

	mencacheObj := mencache.NewMencache(srv.CacheStore)

	services := service.NewService(&service.Deps{
		Ctx:           context.Background(),
		Mencache:      mencacheObj,
		Repositories:  repos,
		Logger:        srv.Logger,
		Observability: traceLoggerObservability,
	})

	handlers := handler.NewHandler(&handler.Deps{
		Service: services,
		Logger:  srv.Logger,
	})

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterCashierQueryServiceServer(gs, handlers.Cashier)
		pb.RegisterCashierCommandServiceServer(gs, handlers.CashierCommand)
		pb.RegisterCashierStatsServiceServer(gs, handlers.CashierStats)
	}

	return srv, nil
}
