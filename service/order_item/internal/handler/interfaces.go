package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order_item"

type OrderItemHandlerGrpc interface {
	pb.OrderItemServiceServer
}
