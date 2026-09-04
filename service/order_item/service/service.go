package service

import (
	"context"

	mencache "github.com/MamangRust/monolith-graphql-pointofsale-order-item/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-order-item/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
)

type Service struct {
	OrderItemQuery OrderItemQueryService
}

type Deps struct {
	Ctx           context.Context
	Mencache      mencache.Mencache
	Repositories  *repository.Repositories
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

func NewService(deps *Deps) *Service {
	return &Service{
		OrderItemQuery: NewOrderItemQueryService(&orderItemQueryDeps{
			Cache:         deps.Mencache,
			Repo:          deps.Repositories.OrderItemQuery,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
