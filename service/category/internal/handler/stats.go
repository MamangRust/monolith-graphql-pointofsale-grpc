package handler

import (
	"context"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"
	"github.com/MamangRust/monolith-point-of-sale-category/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-shared/domain/requests"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/category_errors"
	"go.uber.org/zap"
)

type categoryStatsHandleGrpc struct {
	pb.UnimplementedCategoryStatsServiceServer
	categoryStatsService    service.CategoryStatsService
	categoryStatsById       service.CategoryStatsByIdService
	categoryStatsByMerchant service.CategoryStatsByMerchantService
	logger                  logger.LoggerInterface
}

func NewCategoryStatsHandleGrpc(categoryStatsService service.CategoryStatsService, categoryStatsById service.CategoryStatsByIdService, categoryStatsByMerchant service.CategoryStatsByMerchantService, logger logger.LoggerInterface) pb.CategoryStatsServiceServer {
	return &categoryStatsHandleGrpc{
		categoryStatsService: categoryStatsService,
		logger:               logger,
	}
}

func (s *categoryStatsHandleGrpc) FindMonthlyTotalPrices(ctx context.Context, req *pb.FindYearMonthTotalPrices) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	s.logger.Info("FindMonthlyTotalPrices categories called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if month <= 0 || month > 12 {
		return nil, category_errors.ErrGrpcFailedInvalidMonth
	}

	reqService := requests.MonthTotalPrice{
		Year:  year,
		Month: month,
	}

	methods, err := s.categoryStatsService.FindMonthlyTotalPrice(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthlyTotalPrices categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthlyTotalPrices categories success")

	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapResponseCategoryMonthlyTotalPrices(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	s.logger.Info("FindYearlyTotalPrices categories called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.categoryStatsService.FindYearlyTotalPrice(ctx, year)
	if err != nil {
		s.logger.Error("FindYearlyTotalPrices categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearlyTotalPrices categories success")

	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponseCategoryYearlyTotalPrices(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	s.logger.Info("FindMonthlyTotalPricesById categories called", zap.Int32("id", req.GetCategoryId()))

	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if month <= 0 || month > 12 {
		return nil, category_errors.ErrGrpcFailedInvalidMonth
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthTotalPriceCategory{
		Year:       year,
		Month:      month,
		CategoryID: id,
	}

	methods, err := s.categoryStatsById.FindMonthlyTotalPriceById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthlyTotalPricesById categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthlyTotalPricesById categories success")

	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapResponseCategoryMonthlyTotalPricesById(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	s.logger.Info("FindYearlyTotalPricesById categories called", zap.Int32("id", req.GetCategoryId()))

	year := int(req.GetYear())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearTotalPriceCategory{
		Year:       year,
		CategoryID: id,
	}

	methods, err := s.categoryStatsById.FindYearlyTotalPriceById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearlyTotalPricesById categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearlyTotalPricesById categories success")

	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponseCategoryYearlyTotalPricesById(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	s.logger.Info("FindMonthlyTotalPricesByMerchant categories called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if month <= 0 || month > 12 {
		return nil, category_errors.ErrGrpcFailedInvalidMonth
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.MonthTotalPriceMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	methods, err := s.categoryStatsByMerchant.FindMonthlyTotalPriceByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthlyTotalPricesByMerchant categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthlyTotalPricesByMerchant categories success")

	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapResponseCategoryMonthlyTotalPricesByMerchant(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	s.logger.Info("FindYearlyTotalPricesByMerchant categories called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.YearTotalPriceMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.categoryStatsByMerchant.FindYearlyTotalPriceByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearlyTotalPricesByMerchant categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearlyTotalPricesByMerchant categories success")

	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponseCategoryYearlyTotalPricesByMerchant(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryMonthPrice, error) {
	s.logger.Info("FindMonthPrice categories called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.categoryStatsService.FindMonthPrice(ctx, year)
	if err != nil {
		s.logger.Error("FindMonthPrice categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthPrice categories success")

	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Monthly payment methods retrieved successfully",
		Data:    mapResponsesCategoryMonthlyPrices(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryYearPrice, error) {
	s.logger.Info("FindYearPrice categories called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.categoryStatsService.FindYearPrice(ctx, year)
	if err != nil {
		s.logger.Error("FindYearPrice categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearPrice categories success")

	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponsesCategoryYearlyPrices(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryMonthPrice, error) {
	s.logger.Info("FindMonthPriceByMerchant categories called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.MonthPriceMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.categoryStatsByMerchant.FindMonthPriceByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthPriceByMerchant categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthPriceByMerchant categories success")

	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Merchant monthly payment methods retrieved successfully",
		Data:    mapResponsesCategoryMonthlyPricesByMerchant(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryYearPrice, error) {
	s.logger.Info("FindYearPriceByMerchant categories called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.YearPriceMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.categoryStatsByMerchant.FindYearPriceByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearPriceByMerchant categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearPriceByMerchant categories success")

	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Merchant yearly payment methods retrieved successfully",
		Data:    mapResponsesCategoryYearlyPricesByMerchant(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryMonthPrice, error) {
	s.logger.Info("FindMonthPriceById categories called", zap.Int32("id", req.GetCategoryId()))

	year := int(req.GetYear())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthPriceId{
		Year:       year,
		CategoryID: id,
	}

	methods, err := s.categoryStatsById.FindMonthPriceById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthPriceById categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthPriceById categories success")

	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Merchant monthly payment methods retrieved successfully",
		Data:    mapResponsesCategoryMonthlyPricesById(methods),
	}, nil
}

func (s *categoryStatsHandleGrpc) FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryYearPrice, error) {
	s.logger.Info("FindYearPriceById categories called", zap.Int32("id", req.GetCategoryId()))

	year := int(req.GetYear())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}
	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearPriceId{
		Year:       year,
		CategoryID: id,
	}

	methods, err := s.categoryStatsById.FindYearPriceById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearPriceById categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearPriceById categories success")

	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Merchant yearly payment methods retrieved successfully",
		Data:    mapResponsesCategoryYearlyPricesById(methods),
	}, nil
}
