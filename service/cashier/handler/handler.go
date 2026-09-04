package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-cashier/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Cashier        cashier.CashierQueryServiceServer
	CashierCommand cashier.CashierCommandServiceServer
	CashierStats   cashier.CashierStatsServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewCashierHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		Cashier:        h,
		CashierCommand: h,
		CashierStats:   h,
	}
}
