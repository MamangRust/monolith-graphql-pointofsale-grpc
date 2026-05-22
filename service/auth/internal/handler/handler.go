package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/auth"
	"github.com/MamangRust/monolith-point-of-sale-auth/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
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
