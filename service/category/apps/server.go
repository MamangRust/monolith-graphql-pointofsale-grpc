package apps

import (
	"context"

	"github.com/MamangRust/monolith-graphql-pointofsale-category/handler"
	mencache "github.com/MamangRust/monolith-graphql-pointofsale-category/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-category/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-category/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"
	"google.golang.org/grpc"
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
				pb.RegisterCategoryQueryServiceServer(gs, handlers.Category)
		pb.RegisterCategoryCommandServiceServer(gs, handlers.CategoryCommand)
		pb.RegisterCategoryStatsServiceServer(gs, handlers.CategoryStats)
	}

	return srv, nil
}
