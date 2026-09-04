package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	"github.com/MamangRust/monolith-graphql-pointofsale-product/service"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Product pb.ProductQueryServiceServer
	ProductCommand pb.ProductCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewProductHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		Product:        h,
		ProductCommand: h,
	}
}
