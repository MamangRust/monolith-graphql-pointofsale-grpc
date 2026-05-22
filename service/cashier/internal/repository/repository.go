package repository

import (
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"

	userpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
	merchantpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant")

type Repositories struct {
	UserQuery              UserQueryRepository
	MerchantQuery          MerchantQueryRepository
	CashierQuery           CashierQueryRepository
	CashierCommand         CashierCommandRepository
	CashierStats           CashierStatsRepository
	CashierStatsByMerchant CashierStatByMerchantRepository
	CashierStatsById       CashierStatByIdRepository
}

func NewRepositories(
	DB *db.Queries,
	userClient userpb.UserQueryServiceClient,
	merchantClient merchantpb.MerchantQueryServiceClient,
) *Repositories {
	return &Repositories{
		UserQuery:              NewUserQueryRepository(userClient),
		MerchantQuery:          NewMerchantQueryRepository(merchantClient),
		CashierQuery:           NewCashierQueryRepository(DB),
		CashierCommand:         NewCashierCommandRepository(DB),
		CashierStats:           NewCashierStatsRepository(DB),
		CashierStatsByMerchant: NewCashierStatsByMerchantRepository(DB),
		CashierStatsById:       NewCashierStatsByIdRepository(DB),
	}
}
