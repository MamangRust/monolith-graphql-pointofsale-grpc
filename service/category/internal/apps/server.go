package apps

import (
	"context"

	"github.com/MamangRust/monolith-point-of-sale-category/internal/handler"
	mencache "github.com/MamangRust/monolith-point-of-sale-category/internal/redis"
	"github.com/MamangRust/monolith-point-of-sale-category/internal/repository"
	"github.com/MamangRust/monolith-point-of-sale-category/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/server"
	"github.com/MamangRust/monolith-point-of-sale-shared/observability"
	"google.golang.org/grpc"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepositories(srv.DB)
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
		pb.RegisterCategoryCommandServiceServer(gs, handlers.Command)
		pb.RegisterCategoryQueryServiceServer(gs, handlers.Query)
		pb.RegisterCategoryStatsServiceServer(gs, handlers.Stats)
	}

	return srv, nil
}
