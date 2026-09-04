package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"

type OrderHandleGrpc interface {
	pb.OrderQueryServiceServer
	pb.OrderCommandServiceServer
	pb.OrderStatsServiceServer
}
