package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"

type ProductHandleGrpc interface {
	pb.ProductQueryServiceServer
	pb.ProductCommandServiceServer
}
