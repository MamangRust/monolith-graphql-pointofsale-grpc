package repository

import (
	"context"

	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/errors/merchant_errors"

	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
)

type merchantQueryRepository struct {
	client pbmerchant.MerchantQueryServiceClient
}

func NewMerchantQueryRepository(client pbmerchant.MerchantQueryServiceClient) MerchantQueryRepository {
	return &merchantQueryRepository{
		client: client,
	}
}

func parseNullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (r *merchantQueryRepository) FindById(ctx context.Context, merchantID int) (*db.Merchant, error) {
	resp, err := r.client.FindById(ctx, &pbmerchant.FindByIdMerchantRequest{
		Id: int32(merchantID),
	})
	if err != nil {
		return nil, merchant_errors.ErrFindById
	}

	if resp == nil || resp.Data == nil {
		return nil, merchant_errors.ErrFindById
	}

	m := resp.Data
	res := &db.Merchant{
		MerchantID:   m.Id,
		UserID:       m.UserId,
		Name:         m.Name,
		Description:  parseNullableString(m.Description),
		Address:      parseNullableString(m.Address),
		ContactEmail: parseNullableString(m.ContactEmail),
		ContactPhone: parseNullableString(m.ContactPhone),
		Status:       m.Status,
		CreatedAt:    parsePgTimestamp(m.CreatedAt),
		UpdatedAt:    parsePgTimestamp(m.UpdatedAt),
	}

	return res, nil
}
