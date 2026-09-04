package service

import (
	"context"

	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/domain/requests"
)

type OrderItemQueryService interface {
	FindAllOrderItems(ctx context.Context, req *requests.FindAllOrderItems) ([]*db.GetOrderItemsRow, *int, error)
	FindByActive(ctx context.Context, req *requests.FindAllOrderItems) ([]*db.GetOrderItemsActiveRow, *int, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllOrderItems) ([]*db.GetOrderItemsTrashedRow, *int, error)
	FindOrderItemByOrder(ctx context.Context, orderID int) ([]*db.OrderItem, error)
}
