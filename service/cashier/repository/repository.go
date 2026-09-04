package repository

import (
	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pbuser "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"

)

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
	userClient pbuser.UserQueryServiceClient,
	merchantClient pbmerchant.MerchantQueryServiceClient,
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
