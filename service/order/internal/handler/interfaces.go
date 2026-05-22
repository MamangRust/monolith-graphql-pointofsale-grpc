package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"

type OrderQueryHandler interface {
	pb.OrderQueryServiceServer
}

type OrderCommandHandler interface {
	pb.OrderCommandServiceServer
}

type OrderStatsHandler interface {
	pb.OrderStatsServiceServer
}
