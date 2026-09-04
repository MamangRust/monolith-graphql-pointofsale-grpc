package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	"github.com/MamangRust/monolith-graphql-pointofsale-role/service"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/role"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Role pb.RoleQueryServiceServer
	RoleCommand pb.RoleCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewRoleHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		Role:        h,
		RoleCommand: h,
	}
}
