package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
	"github.com/MamangRust/monolith-graphql-pointofsale-user/service"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	User pb.UserQueryServiceServer
	UserCommand pb.UserCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewUserHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		User:        h,
		UserCommand: h,
	}
}
