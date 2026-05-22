package handler

import pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"

type CategoryQueryHandleGrpc interface {
	pb.CategoryQueryServiceClient
}

type CategoryCommandHandleGrpc interface {
	pb.CategoryCommandServiceClient
}

type CategoryStatsHandleGrpc interface {
	pb.CategoryStatsServiceServer
}
