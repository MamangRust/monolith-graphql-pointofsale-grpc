package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-order-item/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	OrderItem pb.OrderItemServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		OrderItem: NewOrderItemHandleGrpc(
			deps.Service.OrderItemQuery,
			deps.Logger,
		),
	}
}
