package handler

import (
	rolepb "github.com/MamangRust/monolith-graphql-pointofsale-pb/role"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-role/internal/service"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	RoleQuery   rolepb.RoleServiceServer
	RoleCommand rolepb.RoleCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		RoleQuery: NewRoleQueryHandleGrpc(
			deps.Service.RoleQuery,
			deps.Logger,
		),
		RoleCommand: NewRoleCommandHandleGrpc(
			deps.Service.RoleCommand,
			deps.Logger,
		),
	}
}


