package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-user/internal/service"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	UserQuery   pb.UserQueryServiceServer
	UserCommand pb.UserCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		UserQuery: NewUserQueryHandleGrpc(
			deps.Service,
			deps.Logger,
		),
		UserCommand: NewUserCommandHandleGrpc(
			deps.Service,
			deps.Logger,
		),
	}
}
