package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
)

type UserQueryGrpc interface {
	pb.UserQueryServiceServer
}

type UserCommandGrpc interface {
	pb.UserCommandServiceServer
}
