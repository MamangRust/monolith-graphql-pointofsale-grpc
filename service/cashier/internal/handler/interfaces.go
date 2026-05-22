package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"

type CashierQueryHandleGrpc interface {
	pb.CashierQueryServiceServer
}

type CashierCommandHandleGrpc interface {
	pb.CashierCommandServiceServer
}

type CashierStatsHandleGrpc interface {
	pb.CashierStatsServiceServer
}

