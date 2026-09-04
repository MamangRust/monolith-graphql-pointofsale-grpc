package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
)

type CashierHandleGrpc interface {
	pb.CashierQueryServiceServer
	pb.CashierCommandServiceServer
	pb.CashierStatsServiceServer
}
