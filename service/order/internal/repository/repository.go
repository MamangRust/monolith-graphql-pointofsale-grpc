package repository

import (
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"

	cashierpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	merchantpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	orderitempb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order_item"
	productpb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
)

type Repositories struct {
	CashierQuery         CashierQueryRepository
	MerchantQuery        MerchantQueryRepository
	ProductQuery         ProductQueryRepository
	ProductCommand       ProductCommandRepository
	OrderQuery           OrderQueryRepository
	OrderCommand         OrderCommandRepository
	OrderItemQuery       OrderItemQueryRepository
	OrderItemCommand     OrderItemCommandRepository
	OrderStats           OrderStatsRepository
	OrderStatsByMerchant OrderStatByMerchantRepository
}

func NewRepositories(
	DB *db.Queries,
	cashierClient cashierpb.CashierQueryServiceClient,
	merchantClient merchantpb.MerchantQueryServiceClient,
	productClient productpb.ProductQueryServiceClient,
	productCommandClient productpb.ProductCommandServiceClient,
	orderItemClient orderitempb.OrderItemServiceClient,
) *Repositories {
	return &Repositories{
		CashierQuery:         NewCashierQueryRepository(cashierClient),
		MerchantQuery:        NewMerchantQueryRepository(merchantClient),
		ProductQuery:         NewProductQueryRepository(productClient),
		ProductCommand:       NewProductCommandRepository(productCommandClient, productClient),
		OrderQuery:           NewOrderQueryRepository(DB),
		OrderCommand:         NewOrderCommandRepository(DB),
		OrderItemQuery:       NewOrderItemQueryRepository(orderItemClient),
		OrderItemCommand:     NewOrderItemCommandRepository(DB),
		OrderStats:           NewOrderStatsRepository(DB),
		OrderStatsByMerchant: NewOrderStatsByMerchantRepository(DB),
	}
}
