package repository

import (
	"context"

	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/product_errors"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
)

type productQueryRepository struct {
	client pb.ProductQueryServiceClient
}

func NewProductQueryRepository(client pb.ProductQueryServiceClient) ProductQueryRepository {
	return &productQueryRepository{
		client: client,
	}
}

func parseNullableInt32(i int32) *int32 {
	return &i
}

func (r *productQueryRepository) FindById(ctx context.Context, product_id int) (*db.Product, error) {
	resp, err := r.client.FindById(ctx, &pb.FindByIdProductRequest{
		Id: int32(product_id),
	})
	if err != nil {
		return nil, product_errors.ErrFindById
	}

	if resp == nil || resp.Data == nil {
		return nil, product_errors.ErrFindById
	}

	p := resp.Data
	res := &db.Product{
		ProductID:    p.Id,
		MerchantID:   p.MerchantId,
		CategoryID:   p.CategoryId,
		Name:         p.Name,
		Description:  parseNullableString(p.Description),
		Price:        p.Price,
		CountInStock: p.CountInStock,
		Brand:        parseNullableString(p.Brand),
		Weight:       parseNullableInt32(p.Weight),
		SlugProduct:  parseNullableString(p.SlugProduct),
		ImageProduct: parseNullableString(p.ImageProduct),
		Barcode:      parseNullableString(p.Barcode),
		CreatedAt:    parsePgTimestamp(p.CreatedAt),
		UpdatedAt:    parsePgTimestamp(p.UpdatedAt),
	}

	return res, nil
}
