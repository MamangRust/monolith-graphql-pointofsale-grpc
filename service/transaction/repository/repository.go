package repository

import (
	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
	pbcashier "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pborder "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"

)

type Repositories struct {
	CashierQuery                 CashierQueryRepository
	MerchantQuery                MerchantQueryRepository
	OrderQuery                   OrderQueryRepository
	OrderItemQuery               OrderItemQueryRepository
	TransactionCommandRepository TransactionCommandRepository
	TransactionQueryRepository   TransactionQueryRepository
	TransactionStatsRepository   TransactionStatsRepository
	TransactionStatsByMerchant   TransactionStatsByMerchantRepository
}

func NewRepositories(
	DB *db.Queries,
	cashierClient pbcashier.CashierQueryServiceClient,
	merchantClient pbmerchant.MerchantQueryServiceClient,
	orderClient pborder.OrderQueryServiceClient,
	orderItemClient pb.OrderItemServiceClient,
) *Repositories {
	return &Repositories{
		CashierQuery:                 NewCashierQueryRepository(cashierClient),
		MerchantQuery:                NewMerchantQueryRepository(merchantClient),
		OrderQuery:                   NewOrderQueryRepository(orderClient),
		OrderItemQuery:               NewOrderItemQueryRepository(orderItemClient),
		TransactionCommandRepository: NewTransactionCommandRepository(DB),
		TransactionQueryRepository:   NewTransactionQueryRepository(DB),
		TransactionStatsRepository:   NewTransactionStatsRepository(DB),
		TransactionStatsByMerchant:   NewTransactionStatsByMerchantRepository(DB),
	}
}
