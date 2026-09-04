package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
)
type OrderItemHandlerGrpc interface {
	pb.OrderItemServiceServer
}
