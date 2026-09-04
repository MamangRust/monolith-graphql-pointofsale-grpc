package repository

import (
	"context"

	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/errors/order_errors"
	"github.com/jackc/pgx/v5/pgtype"

	pborder "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"
)

type orderQueryRepository struct {
	client pborder.OrderQueryServiceClient
}

func NewOrderQueryRepository(client pborder.OrderQueryServiceClient) OrderQueryRepository {
	return &orderQueryRepository{
		client: client,
	}
}

func (r *orderQueryRepository) FindById(ctx context.Context, order_id int) (*db.Order, error) {
	resp, err := r.client.FindById(ctx, &pborder.FindByIdOrderRequest{
		Id: int32(order_id),
	})
	if err != nil {
		return nil, order_errors.ErrFindById
	}

	if resp == nil || resp.Data == nil {
		return nil, order_errors.ErrFindById
	}

	o := resp.Data
	res := &db.Order{
		OrderID:    o.Id,
		MerchantID: o.MerchantId,
		CashierID:  o.CashierId,
		TotalPrice: int64(o.TotalPrice),
		CreatedAt:  parsePgTimestamp(o.CreatedAt),
		UpdatedAt:  parsePgTimestamp(o.UpdatedAt),
		DeletedAt:  pgtype.Timestamp{},
	}

	return res, nil
}
