package handler

import (
	apipb "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func mapPaginationMeta(meta *apipb.PaginationMeta) *apipb.PaginationMeta {
	if meta == nil {
		return nil
	}
	return &apipb.PaginationMeta{
		CurrentPage:  meta.CurrentPage,
		PageSize:     meta.PageSize,
		TotalPages:   meta.TotalPages,
		TotalRecords: meta.TotalRecords,
	}
}

func mapResponseCashier(cashier *db.Cashier) *pb.CashierResponse {
	if cashier == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	if cashier.CreatedAt.Valid {
		createdAtStr = cashier.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if cashier.UpdatedAt.Valid {
		updatedAtStr = cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &pb.CashierResponse{
		Id:         int32(cashier.CashierID),
		MerchantId: int32(cashier.MerchantID),
		Name:       cashier.Name,
		CreatedAt:  createdAtStr,
		UpdatedAt:  updatedAtStr,
	}
}

func mapResponsesCashier(cashiers []*db.GetCashiersRow) []*pb.CashierResponse {
	var mappedCashiers []*pb.CashierResponse
	for _, cashier := range cashiers {
		var createdAtStr, updatedAtStr string
		if cashier.CreatedAt.Valid {
			createdAtStr = cashier.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if cashier.UpdatedAt.Valid {
			updatedAtStr = cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		mappedCashiers = append(mappedCashiers, &pb.CashierResponse{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  createdAtStr,
			UpdatedAt:  updatedAtStr,
		})
	}
	return mappedCashiers
}

func mapResponsesCashierByMerchant(cashiers []*db.GetCashiersByMerchantRow) []*pb.CashierResponse {
	var mappedCashiers []*pb.CashierResponse
	for _, cashier := range cashiers {
		var createdAtStr, updatedAtStr string
		if cashier.CreatedAt.Valid {
			createdAtStr = cashier.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if cashier.UpdatedAt.Valid {
			updatedAtStr = cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		mappedCashiers = append(mappedCashiers, &pb.CashierResponse{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  createdAtStr,
			UpdatedAt:  updatedAtStr,
		})
	}
	return mappedCashiers
}

func mapResponseCashierDeleteAt(cashier *db.Cashier) *pb.CashierResponseDeleteAt {
	if cashier == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	if cashier.CreatedAt.Valid {
		createdAtStr = cashier.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if cashier.UpdatedAt.Valid {
		updatedAtStr = cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var deletedAt *wrapperspb.StringValue
	if cashier.DeletedAt.Valid {
		deletedAt = wrapperspb.String(cashier.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}

	return &pb.CashierResponseDeleteAt{
		Id:         int32(cashier.CashierID),
		MerchantId: int32(cashier.MerchantID),
		Name:       cashier.Name,
		CreatedAt:  createdAtStr,
		UpdatedAt:  updatedAtStr,
		DeletedAt:  deletedAt,
	}
}

func mapResponsesCashierActive(cashiers []*db.GetCashiersActiveRow) []*pb.CashierResponseDeleteAt {
	var mappedCashiers []*pb.CashierResponseDeleteAt
	for _, cashier := range cashiers {
		var createdAtStr, updatedAtStr string
		if cashier.CreatedAt.Valid {
			createdAtStr = cashier.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if cashier.UpdatedAt.Valid {
			updatedAtStr = cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAt *wrapperspb.StringValue
		if cashier.DeletedAt.Valid {
			deletedAt = wrapperspb.String(cashier.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}
		mappedCashiers = append(mappedCashiers, &pb.CashierResponseDeleteAt{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  createdAtStr,
			UpdatedAt:  updatedAtStr,
			DeletedAt:  deletedAt,
		})
	}
	return mappedCashiers
}

func mapResponsesCashierTrashed(cashiers []*db.GetCashiersTrashedRow) []*pb.CashierResponseDeleteAt {
	var mappedCashiers []*pb.CashierResponseDeleteAt
	for _, cashier := range cashiers {
		var createdAtStr, updatedAtStr string
		if cashier.CreatedAt.Valid {
			createdAtStr = cashier.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		if cashier.UpdatedAt.Valid {
			updatedAtStr = cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAt *wrapperspb.StringValue
		if cashier.DeletedAt.Valid {
			deletedAt = wrapperspb.String(cashier.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}
		mappedCashiers = append(mappedCashiers, &pb.CashierResponseDeleteAt{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  createdAtStr,
			UpdatedAt:  updatedAtStr,
			DeletedAt:  deletedAt,
		})
	}
	return mappedCashiers
}

func mapResponseCashierMonthlySale(cashier *db.GetMonthlyCashierRow) *pb.CashierResponseMonthSales {
	if cashier == nil {
		return nil
	}
	return &pb.CashierResponseMonthSales{
		Month:       cashier.Month,
		CashierId:   int32(cashier.CashierID),
		CashierName: cashier.CashierName,
		OrderCount:  int32(cashier.OrderCount),
		TotalSales:  int32(cashier.TotalSales),
	}
}

func mapResponsesCashierMonthlySales(c []*db.GetMonthlyCashierRow) []*pb.CashierResponseMonthSales {
	var cashierRecords []*pb.CashierResponseMonthSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, mapResponseCashierMonthlySale(cashier))
	}
	return cashierRecords
}

func mapResponseCashierMonthlySaleById(cashier *db.GetMonthlyCashierByCashierIdRow) *pb.CashierResponseMonthSales {
	if cashier == nil {
		return nil
	}
	return &pb.CashierResponseMonthSales{
		Month:       cashier.Month,
		CashierId:   int32(cashier.CashierID),
		CashierName: cashier.CashierName,
		OrderCount:  int32(cashier.OrderCount),
		TotalSales:  int32(cashier.TotalSales),
	}
}

func mapResponsesCashierMonthlySalesById(c []*db.GetMonthlyCashierByCashierIdRow) []*pb.CashierResponseMonthSales {
	var cashierRecords []*pb.CashierResponseMonthSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, mapResponseCashierMonthlySaleById(cashier))
	}
	return cashierRecords
}

func mapResponseCashierMonthlySaleByMerchant(cashier *db.GetMonthlyCashierByMerchantRow) *pb.CashierResponseMonthSales {
	if cashier == nil {
		return nil
	}
	return &pb.CashierResponseMonthSales{
		Month:       cashier.Month,
		CashierId:   int32(cashier.CashierID),
		CashierName: cashier.CashierName,
		OrderCount:  int32(cashier.OrderCount),
		TotalSales:  int32(cashier.TotalSales),
	}
}

func mapResponsesCashierMonthlySalesByMerchant(c []*db.GetMonthlyCashierByMerchantRow) []*pb.CashierResponseMonthSales {
	var cashierRecords []*pb.CashierResponseMonthSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, mapResponseCashierMonthlySaleByMerchant(cashier))
	}
	return cashierRecords
}

func mapResponseCashierYearlySale(cashier *db.GetYearlyCashierRow) *pb.CashierResponseYearSales {
	if cashier == nil {
		return nil
	}
	return &pb.CashierResponseYearSales{
		Year:        cashier.Year,
		CashierId:   int32(cashier.CashierID),
		CashierName: cashier.CashierName,
		OrderCount:  int32(cashier.OrderCount),
		TotalSales:  int32(cashier.TotalSales),
	}
}

func mapResponsesCashierYearlySales(c []*db.GetYearlyCashierRow) []*pb.CashierResponseYearSales {
	var cashierRecords []*pb.CashierResponseYearSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, mapResponseCashierYearlySale(cashier))
	}
	return cashierRecords
}

func mapResponseCashierYearlySaleById(cashier *db.GetYearlyCashierByCashierIdRow) *pb.CashierResponseYearSales {
	if cashier == nil {
		return nil
	}
	return &pb.CashierResponseYearSales{
		Year:        cashier.Year,
		CashierId:   int32(cashier.CashierID),
		CashierName: cashier.CashierName,
		OrderCount:  int32(cashier.OrderCount),
		TotalSales:  int32(cashier.TotalSales),
	}
}

func mapResponsesCashierYearlySalesById(c []*db.GetYearlyCashierByCashierIdRow) []*pb.CashierResponseYearSales {
	var cashierRecords []*pb.CashierResponseYearSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, mapResponseCashierYearlySaleById(cashier))
	}
	return cashierRecords
}

func mapResponseCashierYearlySaleByMerchant(cashier *db.GetYearlyCashierByMerchantRow) *pb.CashierResponseYearSales {
	if cashier == nil {
		return nil
	}
	return &pb.CashierResponseYearSales{
		Year:        cashier.Year,
		CashierId:   int32(cashier.CashierID),
		CashierName: cashier.CashierName,
		OrderCount:  int32(cashier.OrderCount),
		TotalSales:  int32(cashier.TotalSales),
	}
}

func mapResponsesCashierYearlySalesByMerchant(c []*db.GetYearlyCashierByMerchantRow) []*pb.CashierResponseYearSales {
	var cashierRecords []*pb.CashierResponseYearSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, mapResponseCashierYearlySaleByMerchant(cashier))
	}
	return cashierRecords
}

func mapResponseCashierMonthlyTotalSales(c []*db.GetMonthlyTotalSalesCashierRow) []*pb.CashierResponseMonthTotalSales {
	var cashierRecords []*pb.CashierResponseMonthTotalSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, &pb.CashierResponseMonthTotalSales{
			Year:       cashier.Year,
			Month:      cashier.Month,
			TotalSales: cashier.TotalSales,
		})
	}
	return cashierRecords
}

func mapResponseCashierMonthlyTotalSalesById(c []*db.GetMonthlyTotalSalesByIdRow) []*pb.CashierResponseMonthTotalSales {
	var cashierRecords []*pb.CashierResponseMonthTotalSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, &pb.CashierResponseMonthTotalSales{
			Year:       cashier.Year,
			Month:      cashier.Month,
			TotalSales: cashier.TotalSales,
		})
	}
	return cashierRecords
}

func mapResponseCashierMonthlyTotalSalesByMerchant(c []*db.GetMonthlyTotalSalesByMerchantRow) []*pb.CashierResponseMonthTotalSales {
	var cashierRecords []*pb.CashierResponseMonthTotalSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, &pb.CashierResponseMonthTotalSales{
			Year:       cashier.Year,
			Month:      cashier.Month,
			TotalSales: cashier.TotalSales,
		})
	}
	return cashierRecords
}

func mapResponseCashierYearlyTotalSales(c []*db.GetYearlyTotalSalesCashierRow) []*pb.CashierResponseYearTotalSales {
	var cashierRecords []*pb.CashierResponseYearTotalSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, &pb.CashierResponseYearTotalSales{
			Year:       cashier.Year,
			TotalSales: cashier.TotalSales,
		})
	}
	return cashierRecords
}

func mapResponseCashierYearlyTotalSalesById(c []*db.GetYearlyTotalSalesByIdRow) []*pb.CashierResponseYearTotalSales {
	var cashierRecords []*pb.CashierResponseYearTotalSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, &pb.CashierResponseYearTotalSales{
			Year:       cashier.Year,
			TotalSales: cashier.TotalSales,
		})
	}
	return cashierRecords
}

func mapResponseCashierYearlyTotalSalesByMerchant(c []*db.GetYearlyTotalSalesByMerchantRow) []*pb.CashierResponseYearTotalSales {
	var cashierRecords []*pb.CashierResponseYearTotalSales
	for _, cashier := range c {
		cashierRecords = append(cashierRecords, &pb.CashierResponseYearTotalSales{
			Year:       cashier.Year,
			TotalSales: cashier.TotalSales,
		})
	}
	return cashierRecords
}
