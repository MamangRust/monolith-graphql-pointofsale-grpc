package apps

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/hash"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/server"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
	"github.com/MamangRust/monolith-graphql-pointofsale-user/handler"
	mencache "github.com/MamangRust/monolith-graphql-pointofsale-user/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-user/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-user/service"
	"google.golang.org/grpc"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepositories(srv.DB)
	hash := hash.NewHashingPassword()
	traceLoggerObservability := observability.NewTraceLoggerObservability(srv.Logger)

	cacheMetrics, err := observability.NewCacheMetrics("user")
	if err != nil {
		return nil, err
	}

	cacheStore := cache.NewCacheStore(srv.Redis, srv.Logger, cacheMetrics)
	mencacheObj := mencache.NewMencache(cacheStore)

	services := service.NewService(&service.Deps{
		Mencache:      mencacheObj,
		Repositories:  repos,
		Hash:          hash,
		Logger:        srv.Logger,
		Observability: traceLoggerObservability,
	})

	handlers := handler.NewHandler(&handler.Deps{
		Service: services,
		Logger:  srv.Logger,
	})

	srv.RegisterServices = func(gs *grpc.Server) {
				pb.RegisterUserQueryServiceServer(gs, handlers.User)
		pb.RegisterUserCommandServiceServer(gs, handlers.UserCommand)
	}

	return srv, nil
}
