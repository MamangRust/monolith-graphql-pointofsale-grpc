package handler

import (
	pbutils "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Map helpers
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

func mapResponseOrder(order *db.Order) *pb.OrderResponse {
	if order == nil {
		return nil
	}
	return &pb.OrderResponse{
		Id:         int32(order.OrderID),
		MerchantId: int32(order.MerchantID),
		CashierId:  int32(order.CashierID),
		TotalPrice: int32(order.TotalPrice),
		CreatedAt:  order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:  order.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}

func mapResponsesOrder(orders []*db.GetOrdersRow) []*pb.OrderResponse {
	var mappedOrders []*pb.OrderResponse
	for _, order := range orders {
		if order == nil {
			continue
		}
		mappedOrders = append(mappedOrders, &pb.OrderResponse{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt:  order.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return mappedOrders
}

func mapResponseOrderDeleteAt(order *db.Order) *pb.OrderResponseDeleteAt {
	if order == nil {
		return nil
	}
	var deletedAt *wrapperspb.StringValue
	if order.DeletedAt.Valid {
		deletedAt = wrapperspb.String(order.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}

	return &pb.OrderResponseDeleteAt{
		Id:         int32(order.OrderID),
		MerchantId: int32(order.MerchantID),
		CashierId:  int32(order.CashierID),
		TotalPrice: int32(order.TotalPrice),
		CreatedAt:  order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:  order.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}

func mapResponsesOrderActive(orders []*db.GetOrdersActiveRow) []*pb.OrderResponseDeleteAt {
	var mappedOrders []*pb.OrderResponseDeleteAt
	for _, order := range orders {
		if order == nil {
			continue
		}
		var deletedAt *wrapperspb.StringValue
		if order.DeletedAt.Valid {
			deletedAt = wrapperspb.String(order.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}
		mappedOrders = append(mappedOrders, &pb.OrderResponseDeleteAt{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt:  order.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
			DeletedAt:  deletedAt,
		})
	}
	return mappedOrders
}

func mapResponsesOrderTrashed(orders []*db.GetOrdersTrashedRow) []*pb.OrderResponseDeleteAt {
	var mappedOrders []*pb.OrderResponseDeleteAt
	for _, order := range orders {
		if order == nil {
			continue
		}
		var deletedAt *wrapperspb.StringValue
		if order.DeletedAt.Valid {
			deletedAt = wrapperspb.String(order.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}
		mappedOrders = append(mappedOrders, &pb.OrderResponseDeleteAt{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt:  order.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
			DeletedAt:  deletedAt,
		})
	}
	return mappedOrders
}

func mapResponseOrderMonthlyTotalRevenue(row *db.GetMonthlyTotalRevenueRow) *pb.OrderMonthlyTotalRevenueResponse {
	if row == nil {
		return nil
	}
	return &pb.OrderMonthlyTotalRevenueResponse{
		Year:           row.Year,
		Month:          row.Month,
		TotalRevenue:   int32(row.TotalRevenue),
		TotalItemsSold: 0,
	}
}

func mapResponseOrderMonthlyTotalRevenues(c []*db.GetMonthlyTotalRevenueRow) []*pb.OrderMonthlyTotalRevenueResponse {
	var orderRecords []*pb.OrderMonthlyTotalRevenueResponse
	for _, row := range c {
		orderRecords = append(orderRecords, mapResponseOrderMonthlyTotalRevenue(row))
	}
	return orderRecords
}

func mapResponseOrderMonthlyTotalRevenueByMerchant(row *db.GetMonthlyTotalRevenueByMerchantRow) *pb.OrderMonthlyTotalRevenueResponse {
	if row == nil {
		return nil
	}
	return &pb.OrderMonthlyTotalRevenueResponse{
		Year:           row.Year,
		Month:          row.Month,
		TotalRevenue:   int32(row.TotalRevenue),
		TotalItemsSold: 0,
	}
}

func mapResponseOrderMonthlyTotalRevenuesByMerchant(c []*db.GetMonthlyTotalRevenueByMerchantRow) []*pb.OrderMonthlyTotalRevenueResponse {
	var orderRecords []*pb.OrderMonthlyTotalRevenueResponse
	for _, row := range c {
		orderRecords = append(orderRecords, mapResponseOrderMonthlyTotalRevenueByMerchant(row))
	}
	return orderRecords
}

func mapResponseOrderYearlyTotalRevenue(row *db.GetYearlyTotalRevenueRow) *pb.OrderYearlyTotalRevenueResponse {
	if row == nil {
		return nil
	}
	return &pb.OrderYearlyTotalRevenueResponse{
		Year:         row.Year,
		TotalRevenue: int32(row.TotalRevenue),
	}
}

func mapResponseOrderYearlyTotalRevenues(c []*db.GetYearlyTotalRevenueRow) []*pb.OrderYearlyTotalRevenueResponse {
	var orderRecords []*pb.OrderYearlyTotalRevenueResponse
	for _, row := range c {
		orderRecords = append(orderRecords, mapResponseOrderYearlyTotalRevenue(row))
	}
	return orderRecords
}

func mapResponseOrderYearlyTotalRevenueByMerchant(row *db.GetYearlyTotalRevenueByMerchantRow) *pb.OrderYearlyTotalRevenueResponse {
	if row == nil {
		return nil
	}
	return &pb.OrderYearlyTotalRevenueResponse{
		Year:         row.Year,
		TotalRevenue: int32(row.TotalRevenue),
	}
}

func mapResponseOrderYearlyTotalRevenuesByMerchant(c []*db.GetYearlyTotalRevenueByMerchantRow) []*pb.OrderYearlyTotalRevenueResponse {
	var orderRecords []*pb.OrderYearlyTotalRevenueResponse
	for _, row := range c {
		orderRecords = append(orderRecords, mapResponseOrderYearlyTotalRevenueByMerchant(row))
	}
	return orderRecords
}

func mapResponsesOrderMonthlyPrices(c []*db.GetMonthlyOrderRow) []*pb.OrderMonthlyResponse {
	var categoryRecords []*pb.OrderMonthlyResponse
	for _, category := range c {
		if category == nil {
			continue
		}
		categoryRecords = append(categoryRecords, &pb.OrderMonthlyResponse{
			Month:          category.Month,
			OrderCount:     int32(category.OrderCount),
			TotalRevenue:   int32(category.TotalRevenue),
			TotalItemsSold: int32(category.TotalItemsSold),
		})
	}
	return categoryRecords
}

func mapResponsesOrderMonthlyPricesByMerchant(c []*db.GetMonthlyOrderByMerchantRow) []*pb.OrderMonthlyResponse {
	var categoryRecords []*pb.OrderMonthlyResponse
	for _, category := range c {
		if category == nil {
			continue
		}
		categoryRecords = append(categoryRecords, &pb.OrderMonthlyResponse{
			Month:          category.Month,
			OrderCount:     int32(category.OrderCount),
			TotalRevenue:   int32(category.TotalRevenue),
			TotalItemsSold: int32(category.TotalItemsSold),
		})
	}
	return categoryRecords
}

func mapResponsesOrderYearlyPrices(c []*db.GetYearlyOrderRow) []*pb.OrderYearlyResponse {
	var categoryRecords []*pb.OrderYearlyResponse
	for _, category := range c {
		if category == nil {
			continue
		}
		categoryRecords = append(categoryRecords, &pb.OrderYearlyResponse{
			Year:               category.Year,
			OrderCount:         int32(category.OrderCount),
			TotalRevenue:       int32(category.TotalRevenue),
			TotalItemsSold:     int32(category.TotalItemsSold),
			ActiveCashiers:     int32(category.ActiveCashiers),
			UniqueProductsSold: int32(category.UniqueProductsSold),
		})
	}
	return categoryRecords
}

func mapResponsesOrderYearlyPricesByMerchant(c []*db.GetYearlyOrderByMerchantRow) []*pb.OrderYearlyResponse {
	var categoryRecords []*pb.OrderYearlyResponse
	for _, category := range c {
		if category == nil {
			continue
		}
		categoryRecords = append(categoryRecords, &pb.OrderYearlyResponse{
			Year:               category.Year,
			OrderCount:         int32(category.OrderCount),
			TotalRevenue:       int32(category.TotalRevenue),
			TotalItemsSold:     int32(category.TotalItemsSold),
			ActiveCashiers:     int32(category.ActiveCashiers),
			UniqueProductsSold: int32(category.UniqueProductsSold),
		})
	}
	return categoryRecords
}
