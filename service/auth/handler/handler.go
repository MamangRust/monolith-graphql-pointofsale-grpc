package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-auth/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Auth pb.AuthServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		Auth: NewAuthHandleGrpc(
			deps.Service,
			deps.Logger,
		),
	}
}
