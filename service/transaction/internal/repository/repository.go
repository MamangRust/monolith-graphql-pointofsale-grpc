package repository

import (
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"

	cashierpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	merchantpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	orderpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"
	orderitempb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order_item"
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
	cashierClient cashierpb.CashierQueryServiceClient,
	merchantClient merchantpb.MerchantQueryServiceClient,
	orderClient orderpb.OrderQueryServiceClient,
	orderItemClient orderitempb.OrderItemServiceClient,
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
