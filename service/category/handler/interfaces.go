package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"

type CategoryHandleGrpc interface {
	pb.CategoryQueryServiceServer
	pb.CategoryCommandServiceServer
	pb.CategoryStatsServiceServer
}
