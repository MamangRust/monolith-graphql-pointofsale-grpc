package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-category/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Category pb.CategoryQueryServiceServer
	CategoryCommand pb.CategoryCommandServiceServer
	CategoryStats pb.CategoryStatsServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewCategoryHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		Category:        h,
		CategoryCommand: h,
		CategoryStats:   h,
	}
}
