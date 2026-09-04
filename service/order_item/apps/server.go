package apps

import (
	"context"

	"github.com/MamangRust/monolith-graphql-pointofsale-order-item/handler"
	mencache "github.com/MamangRust/monolith-graphql-pointofsale-order-item/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-order-item/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-order-item/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	"google.golang.org/grpc"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepositories(srv.DB)
	traceLoggerObservability := observability.NewTraceLoggerObservability(srv.Logger)

	cacheMetrics, err := observability.NewCacheMetrics("order-item")
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
		pb.RegisterOrderItemServiceServer(gs, handlers.OrderItem)
	}

	return srv, nil
}
