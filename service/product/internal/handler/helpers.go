package handler

import (
	pbutils "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Map helpers
func parseStrPointer(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func parseInt32Pointer(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

func mapPaginationMeta(meta *pbutils.PaginationMeta) *pbutils.PaginationMeta {
	if meta == nil {
		return nil
	}
	return &pbutils.PaginationMeta{
		CurrentPage:  meta.CurrentPage,
		PageSize:     meta.PageSize,
		TotalPages:   meta.TotalPages,
		TotalRecords: meta.TotalRecords,
	}
}

func mapResponseProduct(product *db.Product) *pb.ProductResponse {
	if product == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	if product.CreatedAt.Valid {
		createdAtStr = product.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if product.UpdatedAt.Valid {
		updatedAtStr = product.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &pb.ProductResponse{
		Id:           int32(product.ProductID),
		MerchantId:   int32(product.MerchantID),
		CategoryId:   int32(product.CategoryID),
		Name:         product.Name,
		Description:  parseStrPointer(product.Description),
		Price:        int32(product.Price),
		CountInStock: int32(product.CountInStock),
		Brand:        parseStrPointer(product.Brand),
		Weight:       parseInt32Pointer(product.Weight),
		SlugProduct:  parseStrPointer(product.SlugProduct),
		ImageProduct: parseStrPointer(product.ImageProduct),
		Barcode:      parseStrPointer(product.Barcode),
		CreatedAt:    createdAtStr,
		UpdatedAt:    updatedAtStr,
	}
}

func mapResponsesProduct(products []*db.GetProductsRow) []*pb.ProductResponse {
	var mappedProducts []*pb.ProductResponse
	for _, p := range products {
		if p == nil {
			continue
		}
		var createdAtStr, updatedAtStr string
		if p.CreatedAt.Valid {
			createdAtStr = p.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if p.UpdatedAt.Valid {
			updatedAtStr = p.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		mappedProducts = append(mappedProducts, &pb.ProductResponse{
			Id:           int32(p.ProductID),
			MerchantId:   int32(p.MerchantID),
			CategoryId:   int32(p.CategoryID),
			Name:         p.Name,
			Description:  parseStrPointer(p.Description),
			Price:        int32(p.Price),
			CountInStock: int32(p.CountInStock),
			Brand:        parseStrPointer(p.Brand),
			Weight:       parseInt32Pointer(p.Weight),
			SlugProduct:  parseStrPointer(p.SlugProduct),
			ImageProduct: parseStrPointer(p.ImageProduct),
			Barcode:      parseStrPointer(p.Barcode),
			CreatedAt:    createdAtStr,
			UpdatedAt:    updatedAtStr,
		})
	}
	return mappedProducts
}

func mapResponsesProductByMerchant(products []*db.GetProductsByMerchantRow, merchantId int32) []*pb.ProductResponse {
	var mappedProducts []*pb.ProductResponse
	for _, p := range products {
		if p == nil {
			continue
		}
		var createdAtStr string
		if p.CreatedAt.Valid {
			createdAtStr = p.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		mappedProducts = append(mappedProducts, &pb.ProductResponse{
			Id:           int32(p.ProductID),
			MerchantId:   merchantId,
			Name:         p.Name,
			Description:  parseStrPointer(p.Description),
			Price:        int32(p.Price),
			CountInStock: int32(p.CountInStock),
			Brand:        parseStrPointer(p.Brand),
			ImageProduct: parseStrPointer(p.ImageProduct),
			CreatedAt:    createdAtStr,
		})
	}
	return mappedProducts
}

func mapResponsesProductByCategory(products []*db.GetProductsByCategoryNameRow) []*pb.ProductResponse {
	var mappedProducts []*pb.ProductResponse
	for _, p := range products {
		if p == nil {
			continue
		}
		var createdAtStr, updatedAtStr string
		if p.CreatedAt.Valid {
			createdAtStr = p.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if p.UpdatedAt.Valid {
			updatedAtStr = p.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		mappedProducts = append(mappedProducts, &pb.ProductResponse{
			Id:           int32(p.ProductID),
			MerchantId:   int32(p.MerchantID),
			CategoryId:   int32(p.CategoryID),
			Name:         p.Name,
			Description:  parseStrPointer(p.Description),
			Price:        int32(p.Price),
			CountInStock: int32(p.CountInStock),
			Brand:        parseStrPointer(p.Brand),
			Weight:       parseInt32Pointer(p.Weight),
			SlugProduct:  parseStrPointer(p.SlugProduct),
			ImageProduct: parseStrPointer(p.ImageProduct),
			Barcode:      parseStrPointer(p.Barcode),
			CreatedAt:    createdAtStr,
			UpdatedAt:    updatedAtStr,
		})
	}
	return mappedProducts
}

func mapResponseProductDeleteAt(product *db.Product) *pb.ProductResponseDeleteAt {
	if product == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	if product.CreatedAt.Valid {
		createdAtStr = product.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if product.UpdatedAt.Valid {
		updatedAtStr = product.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var deletedAt *wrapperspb.StringValue
	if product.DeletedAt.Valid {
		deletedAt = wrapperspb.String(product.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}

	return &pb.ProductResponseDeleteAt{
		Id:           int32(product.ProductID),
		MerchantId:   int32(product.MerchantID),
		CategoryId:   int32(product.CategoryID),
		Name:         product.Name,
		Description:  parseStrPointer(product.Description),
		Price:        int32(product.Price),
		CountInStock: int32(product.CountInStock),
		Brand:        parseStrPointer(product.Brand),
		Weight:       parseInt32Pointer(product.Weight),
		SlugProduct:  parseStrPointer(product.SlugProduct),
		ImageProduct: parseStrPointer(product.ImageProduct),
		Barcode:      parseStrPointer(product.Barcode),
		CreatedAt:    createdAtStr,
		UpdatedAt:    updatedAtStr,
		DeletedAt:    deletedAt,
	}
}

func mapResponsesProductActive(products []*db.GetProductsActiveRow) []*pb.ProductResponseDeleteAt {
	var mappedProducts []*pb.ProductResponseDeleteAt
	for _, p := range products {
		if p == nil {
			continue
		}
		var createdAtStr, updatedAtStr string
		if p.CreatedAt.Valid {
			createdAtStr = p.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if p.UpdatedAt.Valid {
			updatedAtStr = p.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAt *wrapperspb.StringValue
		if p.DeletedAt.Valid {
			deletedAt = wrapperspb.String(p.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}

		mappedProducts = append(mappedProducts, &pb.ProductResponseDeleteAt{
			Id:           int32(p.ProductID),
			MerchantId:   int32(p.MerchantID),
			CategoryId:   int32(p.CategoryID),
			Name:         p.Name,
			Description:  parseStrPointer(p.Description),
			Price:        int32(p.Price),
			CountInStock: int32(p.CountInStock),
			Brand:        parseStrPointer(p.Brand),
			Weight:       parseInt32Pointer(p.Weight),
			SlugProduct:  parseStrPointer(p.SlugProduct),
			ImageProduct: parseStrPointer(p.ImageProduct),
			Barcode:      parseStrPointer(p.Barcode),
			CreatedAt:    createdAtStr,
			UpdatedAt:    updatedAtStr,
			DeletedAt:    deletedAt,
		})
	}
	return mappedProducts
}

func mapResponsesProductTrashed(products []*db.GetProductsTrashedRow) []*pb.ProductResponseDeleteAt {
	var mappedProducts []*pb.ProductResponseDeleteAt
	for _, p := range products {
		if p == nil {
			continue
		}
		var createdAtStr, updatedAtStr string
		if p.CreatedAt.Valid {
			createdAtStr = p.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if p.UpdatedAt.Valid {
			updatedAtStr = p.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAt *wrapperspb.StringValue
		if p.DeletedAt.Valid {
			deletedAt = wrapperspb.String(p.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}

		mappedProducts = append(mappedProducts, &pb.ProductResponseDeleteAt{
			Id:           int32(p.ProductID),
			MerchantId:   int32(p.MerchantID),
			CategoryId:   int32(p.CategoryID),
			Name:         p.Name,
			Description:  parseStrPointer(p.Description),
			Price:        int32(p.Price),
			CountInStock: int32(p.CountInStock),
			Brand:        parseStrPointer(p.Brand),
			Weight:       parseInt32Pointer(p.Weight),
			SlugProduct:  parseStrPointer(p.SlugProduct),
			ImageProduct: parseStrPointer(p.ImageProduct),
			Barcode:      parseStrPointer(p.Barcode),
			CreatedAt:    createdAtStr,
			UpdatedAt:    updatedAtStr,
			DeletedAt:    deletedAt,
		})
	}
	return mappedProducts
}
