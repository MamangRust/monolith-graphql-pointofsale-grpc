package handler

import (
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pbcat "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"
)

// Internal map helpers
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

func mapResponseCategory(category *db.Category) *pbcat.CategoryResponse {
	if category == nil {
		return nil
	}
	var description, slugCategory string
	if category.Description != nil {
		description = *category.Description
	}
	if category.SlugCategory != nil {
		slugCategory = *category.SlugCategory
	}
	var createdAtStr, updatedAtStr string
	if category.CreatedAt.Valid {
		createdAtStr = category.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if category.UpdatedAt.Valid {
		updatedAtStr = category.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &pbcat.CategoryResponse{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  description,
		SlugCategory: slugCategory,
		CreatedAt:    createdAtStr,
		UpdatedAt:    updatedAtStr,
	}
}

func mapResponseGetCategory(category *db.GetCategoriesRow) *pbcat.CategoryResponse {
	if category == nil {
		return nil
	}
	var description, slugCategory string
	if category.Description != nil {
		description = *category.Description
	}
	if category.SlugCategory != nil {
		slugCategory = *category.SlugCategory
	}
	var createdAtStr, updatedAtStr string
	if category.CreatedAt.Valid {
		createdAtStr = category.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if category.UpdatedAt.Valid {
		updatedAtStr = category.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &pbcat.CategoryResponse{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  description,
		SlugCategory: slugCategory,
		CreatedAt:    createdAtStr,
		UpdatedAt:    updatedAtStr,
	}
}

func mapResponsesCategory(categories []*db.GetCategoriesRow) []*pbcat.CategoryResponse {
	var mappedCategories []*pbcat.CategoryResponse
	for _, category := range categories {
		mappedCategories = append(mappedCategories, mapResponseGetCategory(category))
	}
	return mappedCategories
}

func mapResponseCategoryDeleteAt(category *db.Category) *pbcat.CategoryResponseDeleteAt {
	if category == nil {
		return nil
	}
	var description, slugCategory string
	if category.Description != nil {
		description = *category.Description
	}
	if category.SlugCategory != nil {
		slugCategory = *category.SlugCategory
	}
	var createdAtStr, updatedAtStr string
	if category.CreatedAt.Valid {
		createdAtStr = category.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if category.UpdatedAt.Valid {
		updatedAtStr = category.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var deletedAt *wrapperspb.StringValue
	if category.DeletedAt.Valid {
		deletedAt = wrapperspb.String(category.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}

	return &pbcat.CategoryResponseDeleteAt{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  description,
		SlugCategory: slugCategory,
		CreatedAt:    createdAtStr,
		UpdatedAt:    updatedAtStr,
		DeletedAt:    deletedAt,
	}
}

func mapResponseGetCategoryActive(category *db.GetCategoriesActiveRow) *pbcat.CategoryResponseDeleteAt {
	if category == nil {
		return nil
	}
	var description, slugCategory string
	if category.Description != nil {
		description = *category.Description
	}
	if category.SlugCategory != nil {
		slugCategory = *category.SlugCategory
	}
	var createdAtStr, updatedAtStr string
	if category.CreatedAt.Valid {
		createdAtStr = category.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if category.UpdatedAt.Valid {
		updatedAtStr = category.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var deletedAt *wrapperspb.StringValue
	if category.DeletedAt.Valid {
		deletedAt = wrapperspb.String(category.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}

	return &pbcat.CategoryResponseDeleteAt{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  description,
		SlugCategory: slugCategory,
		CreatedAt:    createdAtStr,
		UpdatedAt:    updatedAtStr,
		DeletedAt:    deletedAt,
	}
}

func mapResponsesCategoryActive(categories []*db.GetCategoriesActiveRow) []*pbcat.CategoryResponseDeleteAt {
	var mappedCategories []*pbcat.CategoryResponseDeleteAt
	for _, category := range categories {
		mappedCategories = append(mappedCategories, mapResponseGetCategoryActive(category))
	}
	return mappedCategories
}

func mapResponseGetCategoryTrashed(category *db.GetCategoriesTrashedRow) *pbcat.CategoryResponseDeleteAt {
	if category == nil {
		return nil
	}
	var description, slugCategory string
	if category.Description != nil {
		description = *category.Description
	}
	if category.SlugCategory != nil {
		slugCategory = *category.SlugCategory
	}
	var createdAtStr, updatedAtStr string
	if category.CreatedAt.Valid {
		createdAtStr = category.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if category.UpdatedAt.Valid {
		updatedAtStr = category.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var deletedAt *wrapperspb.StringValue
	if category.DeletedAt.Valid {
		deletedAt = wrapperspb.String(category.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}

	return &pbcat.CategoryResponseDeleteAt{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  description,
		SlugCategory: slugCategory,
		CreatedAt:    createdAtStr,
		UpdatedAt:    updatedAtStr,
		DeletedAt:    deletedAt,
	}
}

func mapResponsesCategoryTrashed(categories []*db.GetCategoriesTrashedRow) []*pbcat.CategoryResponseDeleteAt {
	var mappedCategories []*pbcat.CategoryResponseDeleteAt
	for _, category := range categories {
		mappedCategories = append(mappedCategories, mapResponseGetCategoryTrashed(category))
	}
	return mappedCategories
}

func mapResponseCategoryMonthlyPrice(category *db.GetMonthlyCategoryRow) *pbcat.CategoryMonthPriceResponse {
	if category == nil {
		return nil
	}
	return &pbcat.CategoryMonthPriceResponse{
		Month:        category.Month,
		CategoryId:   category.CategoryID,
		CategoryName: category.CategoryName,
		OrderCount:   int32(category.OrderCount),
		ItemsSold:    int32(category.ItemsSold),
		TotalRevenue: category.TotalRevenue,
	}
}

func mapResponsesCategoryMonthlyPrices(c []*db.GetMonthlyCategoryRow) []*pbcat.CategoryMonthPriceResponse {
	var categoryRecords []*pbcat.CategoryMonthPriceResponse
	for _, category := range c {
		categoryRecords = append(categoryRecords, mapResponseCategoryMonthlyPrice(category))
	}
	return categoryRecords
}

func mapResponseCategoryMonthlyPriceById(category *db.GetMonthlyCategoryByIdRow) *pbcat.CategoryMonthPriceResponse {
	if category == nil {
		return nil
	}
	return &pbcat.CategoryMonthPriceResponse{
		Month:        category.Month,
		CategoryId:   category.CategoryID,
		CategoryName: category.CategoryName,
		OrderCount:   int32(category.OrderCount),
		ItemsSold:    int32(category.ItemsSold),
		TotalRevenue: category.TotalRevenue,
	}
}

func mapResponsesCategoryMonthlyPricesById(c []*db.GetMonthlyCategoryByIdRow) []*pbcat.CategoryMonthPriceResponse {
	var categoryRecords []*pbcat.CategoryMonthPriceResponse
	for _, category := range c {
		categoryRecords = append(categoryRecords, mapResponseCategoryMonthlyPriceById(category))
	}
	return categoryRecords
}

func mapResponseCategoryMonthlyPriceByMerchant(category *db.GetMonthlyCategoryByMerchantRow) *pbcat.CategoryMonthPriceResponse {
	if category == nil {
		return nil
	}
	return &pbcat.CategoryMonthPriceResponse{
		Month:        category.Month,
		CategoryId:   category.CategoryID,
		CategoryName: category.CategoryName,
		OrderCount:   int32(category.OrderCount),
		ItemsSold:    int32(category.ItemsSold),
		TotalRevenue: category.TotalRevenue,
	}
}

func mapResponsesCategoryMonthlyPricesByMerchant(c []*db.GetMonthlyCategoryByMerchantRow) []*pbcat.CategoryMonthPriceResponse {
	var categoryRecords []*pbcat.CategoryMonthPriceResponse
	for _, category := range c {
		categoryRecords = append(categoryRecords, mapResponseCategoryMonthlyPriceByMerchant(category))
	}
	return categoryRecords
}

func mapResponseCategoryYearlyPrice(category *db.GetYearlyCategoryRow) *pbcat.CategoryYearPriceResponse {
	if category == nil {
		return nil
	}
	return &pbcat.CategoryYearPriceResponse{
		Year:               category.Year,
		CategoryId:         category.CategoryID,
		CategoryName:       category.CategoryName,
		OrderCount:         int32(category.OrderCount),
		ItemsSold:          int32(category.ItemsSold),
		TotalRevenue:       category.TotalRevenue,
		UniqueProductsSold: int32(category.UniqueProductsSold),
	}
}

func mapResponsesCategoryYearlyPrices(c []*db.GetYearlyCategoryRow) []*pbcat.CategoryYearPriceResponse {
	var categoryRecords []*pbcat.CategoryYearPriceResponse
	for _, category := range c {
		categoryRecords = append(categoryRecords, mapResponseCategoryYearlyPrice(category))
	}
	return categoryRecords
}

func mapResponseCategoryYearlyPriceById(category *db.GetYearlyCategoryByIdRow) *pbcat.CategoryYearPriceResponse {
	if category == nil {
		return nil
	}
	return &pbcat.CategoryYearPriceResponse{
		Year:               category.Year,
		CategoryId:         category.CategoryID,
		CategoryName:       category.CategoryName,
		OrderCount:         int32(category.OrderCount),
		ItemsSold:          int32(category.ItemsSold),
		TotalRevenue:       category.TotalRevenue,
		UniqueProductsSold: int32(category.UniqueProductsSold),
	}
}

func mapResponsesCategoryYearlyPricesById(c []*db.GetYearlyCategoryByIdRow) []*pbcat.CategoryYearPriceResponse {
	var categoryRecords []*pbcat.CategoryYearPriceResponse
	for _, category := range c {
		categoryRecords = append(categoryRecords, mapResponseCategoryYearlyPriceById(category))
	}
	return categoryRecords
}

func mapResponseCategoryYearlyPriceByMerchant(category *db.GetYearlyCategoryByMerchantRow) *pbcat.CategoryYearPriceResponse {
	if category == nil {
		return nil
	}
	return &pbcat.CategoryYearPriceResponse{
		Year:               category.Year,
		CategoryId:         category.CategoryID,
		CategoryName:       category.CategoryName,
		OrderCount:         int32(category.OrderCount),
		ItemsSold:          int32(category.ItemsSold),
		TotalRevenue:       category.TotalRevenue,
		UniqueProductsSold: int32(category.UniqueProductsSold),
	}
}

func mapResponsesCategoryYearlyPricesByMerchant(c []*db.GetYearlyCategoryByMerchantRow) []*pbcat.CategoryYearPriceResponse {
	var categoryRecords []*pbcat.CategoryYearPriceResponse
	for _, category := range c {
		categoryRecords = append(categoryRecords, mapResponseCategoryYearlyPriceByMerchant(category))
	}
	return categoryRecords
}

func mapResponseCashierMonthlyTotalPrice(c *db.GetMonthlyTotalPriceRow) *pbcat.CategoriesMonthlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &pbcat.CategoriesMonthlyTotalPriceResponse{
		Year:         c.Year,
		Month:        c.Month,
		TotalRevenue: c.TotalRevenue,
	}
}

func mapResponseCategoryMonthlyTotalPrices(c []*db.GetMonthlyTotalPriceRow) []*pbcat.CategoriesMonthlyTotalPriceResponse {
	var CategoryRecords []*pbcat.CategoriesMonthlyTotalPriceResponse
	for _, Category := range c {
		CategoryRecords = append(CategoryRecords, mapResponseCashierMonthlyTotalPrice(Category))
	}
	return CategoryRecords
}

func mapResponseCashierMonthlyTotalPriceById(c *db.GetMonthlyTotalPriceByIdRow) *pbcat.CategoriesMonthlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &pbcat.CategoriesMonthlyTotalPriceResponse{
		Year:         c.Year,
		Month:        c.Month,
		TotalRevenue: c.TotalRevenue,
	}
}

func mapResponseCategoryMonthlyTotalPricesById(c []*db.GetMonthlyTotalPriceByIdRow) []*pbcat.CategoriesMonthlyTotalPriceResponse {
	var CategoryRecords []*pbcat.CategoriesMonthlyTotalPriceResponse
	for _, Category := range c {
		CategoryRecords = append(CategoryRecords, mapResponseCashierMonthlyTotalPriceById(Category))
	}
	return CategoryRecords
}

func mapResponseCashierMonthlyTotalPriceByMerchant(c *db.GetMonthlyTotalPriceByMerchantRow) *pbcat.CategoriesMonthlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &pbcat.CategoriesMonthlyTotalPriceResponse{
		Year:         c.Year,
		Month:        c.Month,
		TotalRevenue: c.TotalRevenue,
	}
}

func mapResponseCategoryMonthlyTotalPricesByMerchant(c []*db.GetMonthlyTotalPriceByMerchantRow) []*pbcat.CategoriesMonthlyTotalPriceResponse {
	var CategoryRecords []*pbcat.CategoriesMonthlyTotalPriceResponse
	for _, Category := range c {
		CategoryRecords = append(CategoryRecords, mapResponseCashierMonthlyTotalPriceByMerchant(Category))
	}
	return CategoryRecords
}

func mapResponseCategoryYearlyTotalSale(c *db.GetYearlyTotalPriceRow) *pbcat.CategoriesYearlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &pbcat.CategoriesYearlyTotalPriceResponse{
		Year:         c.Year,
		TotalRevenue: c.TotalRevenue,
	}
}

func mapResponseCategoryYearlyTotalPrices(c []*db.GetYearlyTotalPriceRow) []*pbcat.CategoriesYearlyTotalPriceResponse {
	var CategoryRecords []*pbcat.CategoriesYearlyTotalPriceResponse
	for _, Category := range c {
		CategoryRecords = append(CategoryRecords, mapResponseCategoryYearlyTotalSale(Category))
	}
	return CategoryRecords
}

func mapResponseCategoryYearlyTotalSaleById(c *db.GetYearlyTotalPriceByIdRow) *pbcat.CategoriesYearlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &pbcat.CategoriesYearlyTotalPriceResponse{
		Year:         c.Year,
		TotalRevenue: c.TotalRevenue,
	}
}

func mapResponseCategoryYearlyTotalPricesById(c []*db.GetYearlyTotalPriceByIdRow) []*pbcat.CategoriesYearlyTotalPriceResponse {
	var CategoryRecords []*pbcat.CategoriesYearlyTotalPriceResponse
	for _, Category := range c {
		CategoryRecords = append(CategoryRecords, mapResponseCategoryYearlyTotalSaleById(Category))
	}
	return CategoryRecords
}

func mapResponseCategoryYearlyTotalSaleByMerchant(c *db.GetYearlyTotalPriceByMerchantRow) *pbcat.CategoriesYearlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &pbcat.CategoriesYearlyTotalPriceResponse{
		Year:         c.Year,
		TotalRevenue: c.TotalRevenue,
	}
}

func mapResponseCategoryYearlyTotalPricesByMerchant(c []*db.GetYearlyTotalPriceByMerchantRow) []*pbcat.CategoriesYearlyTotalPriceResponse {
	var CategoryRecords []*pbcat.CategoriesYearlyTotalPriceResponse
	for _, Category := range c {
		CategoryRecords = append(CategoryRecords, mapResponseCategoryYearlyTotalSaleByMerchant(Category))
	}
	return CategoryRecords
}
