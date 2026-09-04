package repository

import (
	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
	pbcashier "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pbproduct "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"

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
	cashierClient pbcashier.CashierQueryServiceClient,
	merchantClient pbmerchant.MerchantQueryServiceClient,
	productClient pbproduct.ProductQueryServiceClient,
	orderItemClient pb.OrderItemServiceClient,
) *Repositories {
	return &Repositories{
		CashierQuery:         NewCashierQueryRepository(cashierClient),
		MerchantQuery:        NewMerchantQueryRepository(merchantClient),
		ProductQuery:         NewProductQueryRepository(productClient),
		ProductCommand:       NewProductCommandRepository(DB),
		OrderQuery:           NewOrderQueryRepository(DB),
		OrderCommand:         NewOrderCommandRepository(DB),
		OrderItemQuery:       NewOrderItemQueryRepository(orderItemClient),
		OrderItemCommand:     NewOrderItemCommandRepository(DB),
		OrderStats:           NewOrderStatsRepository(DB),
		OrderStatsByMerchant: NewOrderStatsByMerchantRepository(DB),
	}
}
