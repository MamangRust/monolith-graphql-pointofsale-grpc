package handler

import (
	"github.com/MamangRust/monolith-point-of-sale-order/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	OrderQuery   OrderQueryHandler
	OrderCommand OrderCommandHandler
	OrderStats   OrderStatsHandler
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		OrderQuery: NewOrderQueryHandler(
			deps.Service,
			deps.Logger,
		),
		OrderCommand: NewOrderCommandHandler(
			deps.Service,
			deps.Logger,
		),
		OrderStats: NewOrderStatsHandleGrpc(
			deps.Service,
			deps.Logger,
		),
	}
}
