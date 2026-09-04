package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/transaction"

type TransactionHandleGrpc interface {
	pb.TransactionQueryServiceServer
	pb.TransactionCommandServiceServer
	pb.TransactionStatsServiceServer
}
