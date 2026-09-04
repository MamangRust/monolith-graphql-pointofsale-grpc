package apps

import (
	"context"

	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-product/handler"
	mencache "github.com/MamangRust/monolith-graphql-pointofsale-product/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-product/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-product/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
	"google.golang.org/grpc"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepositories(srv.DB)

	mencacheObj := mencache.NewMencache(srv.CacheStore)

	obs := observability.NewTraceLoggerObservability(srv.Logger)

	services := service.NewService(&service.Deps{
		Mencache:      mencacheObj,
		Ctx:           context.Background(),
		Repositories:  repos,
		Logger:        srv.Logger,
		Observability: obs,
	})

	handlers := handler.NewHandler(&handler.Deps{
		Service: services,
		Logger:  srv.Logger,
	})

	srv.RegisterServices = func(gs *grpc.Server) {
				pb.RegisterProductQueryServiceServer(gs, handlers.Product)
		pb.RegisterProductCommandServiceServer(gs, handlers.ProductCommand)
	}

	return srv, nil
}
