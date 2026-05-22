package handler

import (
	"github.com/MamangRust/monolith-point-of-sale-cashier/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Query   pb.CashierQueryServiceServer
	Command pb.CashierCommandServiceServer
	Stats   pb.CashierStatsServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		Query: NewCashierQueryHandleGrpc(
			deps.Service.CashierQuery,
			deps.Logger,
		),
		Command: NewCashierCommandHandleGrpc(
			deps.Service.CashierCommand,
			deps.Logger,
		),
		Stats: NewCashierStatsHandleGrpc(
			deps.Service.CashierStats,
			deps.Service.CashierStatsById,
			deps.Service.CashierStatsByMerchant,
			deps.Logger,
		),
	}
}

