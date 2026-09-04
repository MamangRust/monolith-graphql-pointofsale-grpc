package repository

import (
	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
)

type Repositories struct {
	OrderItemQuery OrderItemQueryRepository
}

func NewRepositories(DB *db.Queries) *Repositories {
	return &Repositories{
		OrderItemQuery: NewOrderItemQueryRepository(DB),
	}
}
