package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-order/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Order pb.OrderQueryServiceServer
	OrderCommand pb.OrderCommandServiceServer
	OrderStats pb.OrderStatsServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewOrderHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		Order:        h,
		OrderCommand: h,
		OrderStats:   h,
	}
}
