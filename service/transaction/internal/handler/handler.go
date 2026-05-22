package handler

import (
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-transacton/internal/service"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	TransactionQuery   TransactionQueryHandleGrpc
	TransactionCommand TransactionCommandHandleGrpc
	TransactionStats   TransactionStatsHandleGrpc
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		TransactionQuery: NewTransactionQueryHandleGrpc(
			deps.Service,
			deps.Logger,
		),
		TransactionCommand: NewTransactionCommandHandleGrpc(
			deps.Service.TransactionCommand,
			deps.Logger,
		),
		TransactionStats: NewTransactionStatsHandleGrpc(
			deps.Service.TransactionStats,
			deps.Service.TransactionStatsByMerchant,
			deps.Logger,
		),
	}
}
