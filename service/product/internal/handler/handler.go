package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-product/internal/service"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	ProductQuery   pb.ProductQueryServiceServer
	ProductCommand pb.ProductCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		ProductQuery: NewProductQueryHandleGrpc(
			deps.Service,
			deps.Logger,
		),
		ProductCommand: NewProductCommandHandleGrpc(
			deps.Service,
			deps.Logger,
		),
	}
}
