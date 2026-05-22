package handler

import (
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"github.com/jackc/pgx/v5/pgtype"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
)

func mapSqlNullString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func mapSqlNullTime(t pgtype.Timestamp) string {
	if t.Valid {
		return t.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}

// Map helpers
func mapPaginationMeta(meta *pb.PaginationMeta) *pb.PaginationMeta {
	if meta == nil {
		return nil
	}
	return &pb.PaginationMeta{
		CurrentPage:  meta.CurrentPage,
		PageSize:     meta.PageSize,
		TotalPages:   meta.TotalPages,
		TotalRecords: meta.TotalRecords,
	}
}

func mapResponseMerchant(merchant *db.Merchant) *pbmerchant.MerchantResponse {
	if merchant == nil {
		return nil
	}
	return &pbmerchant.MerchantResponse{
		Id:           int32(merchant.MerchantID),
		UserId:       int32(merchant.UserID),
		Name:         merchant.Name,
		Description:  mapSqlNullString(merchant.Description),
		Address:      mapSqlNullString(merchant.Address),
		ContactEmail: mapSqlNullString(merchant.ContactEmail),
		ContactPhone: mapSqlNullString(merchant.ContactPhone),
		Status:       merchant.Status,
		CreatedAt:    mapSqlNullTime(merchant.CreatedAt),
		UpdatedAt:    mapSqlNullTime(merchant.UpdatedAt),
	}
}

func mapResponsesGetMerchantsRow(merchants []*db.GetMerchantsRow) []*pbmerchant.MerchantResponse {
	var mapped []*pbmerchant.MerchantResponse
	for _, m := range merchants {
		mapped = append(mapped, &pbmerchant.MerchantResponse{
			Id:           int32(m.MerchantID),
			UserId:       int32(m.UserID),
			Name:         m.Name,
			Description:  mapSqlNullString(m.Description),
			Address:      mapSqlNullString(m.Address),
			ContactEmail: mapSqlNullString(m.ContactEmail),
			ContactPhone: mapSqlNullString(m.ContactPhone),
			Status:       m.Status,
			CreatedAt:    mapSqlNullTime(m.CreatedAt),
			UpdatedAt:    mapSqlNullTime(m.UpdatedAt),
		})
	}
	return mapped
}

func mapResponseMerchantDeleteAt(merchant *db.Merchant) *pbmerchant.MerchantResponseDeleteAt {
	if merchant == nil {
		return nil
	}
	return &pbmerchant.MerchantResponseDeleteAt{
		Id:           int32(merchant.MerchantID),
		UserId:       int32(merchant.UserID),
		Name:         merchant.Name,
		Description:  mapSqlNullString(merchant.Description),
		Address:      mapSqlNullString(merchant.Address),
		ContactEmail: mapSqlNullString(merchant.ContactEmail),
		ContactPhone: mapSqlNullString(merchant.ContactPhone),
		Status:       merchant.Status,
		CreatedAt:    mapSqlNullTime(merchant.CreatedAt),
		UpdatedAt:    mapSqlNullTime(merchant.UpdatedAt),
		DeletedAt:    mapSqlNullTime(merchant.DeletedAt),
	}
}

func mapResponsesGetMerchantsActiveRow(merchants []*db.GetMerchantsActiveRow) []*pbmerchant.MerchantResponseDeleteAt {
	var mapped []*pbmerchant.MerchantResponseDeleteAt
	for _, m := range merchants {
		mapped = append(mapped, &pbmerchant.MerchantResponseDeleteAt{
			Id:           int32(m.MerchantID),
			UserId:       int32(m.UserID),
			Name:         m.Name,
			Description:  mapSqlNullString(m.Description),
			Address:      mapSqlNullString(m.Address),
			ContactEmail: mapSqlNullString(m.ContactEmail),
			ContactPhone: mapSqlNullString(m.ContactPhone),
			Status:       m.Status,
			CreatedAt:    mapSqlNullTime(m.CreatedAt),
			UpdatedAt:    mapSqlNullTime(m.UpdatedAt),
			DeletedAt:    mapSqlNullTime(m.DeletedAt),
		})
	}
	return mapped
}

func mapResponsesGetMerchantsTrashedRow(merchants []*db.GetMerchantsTrashedRow) []*pbmerchant.MerchantResponseDeleteAt {
	var mapped []*pbmerchant.MerchantResponseDeleteAt
	for _, m := range merchants {
		mapped = append(mapped, &pbmerchant.MerchantResponseDeleteAt{
			Id:           int32(m.MerchantID),
			UserId:       int32(m.UserID),
			Name:         m.Name,
			Description:  mapSqlNullString(m.Description),
			Address:      mapSqlNullString(m.Address),
			ContactEmail: mapSqlNullString(m.ContactEmail),
			ContactPhone: mapSqlNullString(m.ContactPhone),
			Status:       m.Status,
			CreatedAt:    mapSqlNullTime(m.CreatedAt),
			UpdatedAt:    mapSqlNullTime(m.UpdatedAt),
			DeletedAt:    mapSqlNullTime(m.DeletedAt),
		})
	}
	return mapped
}
