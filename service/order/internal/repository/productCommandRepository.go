package repository

import (
	"context"

	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/product_errors"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
)

type productCommandRepository struct {
	client      pb.ProductCommandServiceClient
	queryClient pb.ProductQueryServiceClient
}

func NewProductCommandRepository(client pb.ProductCommandServiceClient, queryClient pb.ProductQueryServiceClient) ProductCommandRepository {
	return &productCommandRepository{
		client:      client,
		queryClient: queryClient,
	}
}

func (r *productCommandRepository) UpdateProductCountStock(ctx context.Context, productID int, stock int) (*db.Product, error) {
	// 1. Fetch current product first
	getResp, err := r.queryClient.FindById(ctx, &pb.FindByIdProductRequest{
		Id: int32(productID),
	})
	if err != nil {
		return nil, product_errors.ErrUpdateProductCountStock
	}

	if getResp == nil || getResp.Data == nil {
		return nil, product_errors.ErrUpdateProductCountStock
	}

	curr := getResp.Data

	// 2. Perform the update request with new stock value
	updateReq := &pb.UpdateProductRequest{
		ProductId:    curr.Id,
		MerchantId:   curr.MerchantId,
		CategoryId:   curr.CategoryId,
		Name:         curr.Name,
		Description:  curr.Description,
		Price:        curr.Price,
		CountInStock: int32(stock),
		Brand:        curr.Brand,
		Weight:       curr.Weight,
		ImageProduct: curr.ImageProduct,
	}

	updateResp, err := r.client.Update(ctx, updateReq)
	if err != nil {
		return nil, product_errors.ErrUpdateProductCountStock
	}

	if updateResp == nil || updateResp.Data == nil {
		return nil, product_errors.ErrUpdateProductCountStock
	}

	p := updateResp.Data
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
