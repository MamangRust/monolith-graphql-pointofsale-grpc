package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/transaction"
	"github.com/MamangRust/monolith-graphql-pointofsale-transacton/service"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Transaction pb.TransactionQueryServiceServer
	TransactionCommand pb.TransactionCommandServiceServer
	TransactionStats pb.TransactionStatsServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewTransactionHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		Transaction:        h,
		TransactionCommand: h,
		TransactionStats:   h,
	}
}
