package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"
	"github.com/MamangRust/monolith-point-of-sale-category/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Query   pb.CategoryQueryServiceServer
	Command pb.CategoryCommandServiceServer
	Stats   pb.CategoryStatsServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		Query: NewcategoryQueryHandleGrpc(
			deps.Service.CategoryQuery,
			deps.Logger,
		),
		Command: NewcategoryCommandHandleGrpc(
			deps.Service.CategoryCommand,
			deps.Logger,
		),
		Stats: NewCategoryStatsHandleGrpc(
			deps.Service.CategoryStats,
			deps.Service.CategoryStatsById,
			deps.Service.CategoryStatsByMerchant,
			deps.Logger,
		),
	}
}
