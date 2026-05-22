package handler

import (
	"context"

	"github.com/MamangRust/monolith-point-of-sale-cashier/internal/service"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/cashier"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-shared/domain/requests"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/cashier_errors"
	"go.uber.org/zap"
)

type cashierStatsHandleGrpc struct {
	pb.UnimplementedCashierStatsServiceServer
	cashierStats           service.CashierStatsService
	cashierStatsById       service.CashierStatsByIdService
	cashierStatsByMerchant service.CashierStatsByMerchant
	logger                 logger.LoggerInterface
}

func NewCashierStatsHandleGrpc(
	cashierStats service.CashierStatsService,
	cashierStatsById service.CashierStatsByIdService,
	cashierStatsByMerchant service.CashierStatsByMerchant,
	logger logger.LoggerInterface,
) pb.CashierStatsServiceServer {
	return &cashierStatsHandleGrpc{
		cashierStats:           cashierStats,
		cashierStatsById:       cashierStatsById,
		cashierStatsByMerchant: cashierStatsByMerchant,
		logger:                 logger,
	}
}

func (s *cashierStatsHandleGrpc) FindMonthlyTotalSales(ctx context.Context, req *pb.FindYearMonthTotalSales) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	s.logger.Info("FindMonthlyTotalSales cashier called", zap.Int32("year", req.GetYear()), zap.Int32("month", req.GetMonth()))

	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if month <= 0 || month > 12 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMonth
	}

	reqService := requests.MonthTotalSales{
		Year:  year,
		Month: month,
	}

	methods, err := s.cashierStats.FindMonthlyTotalSales(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthlyTotalSales cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthlyTotalSales cashier success")

	return &pb.ApiResponseCashierMonthlyTotalSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapResponseCashierMonthlyTotalSales(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindYearlyTotalSales(ctx context.Context, req *pb.FindYearTotalSales) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	s.logger.Info("FindYearlyTotalSales cashier called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.cashierStats.FindYearlyTotalSales(ctx, year)
	if err != nil {
		s.logger.Error("FindYearlyTotalSales cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearlyTotalSales cashier success")

	return &pb.ApiResponseCashierYearlyTotalSales{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponseCashierYearlyTotalSales(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindMonthlyTotalSalesById(ctx context.Context, req *pb.FindYearMonthTotalSalesById) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	s.logger.Info("FindMonthlyTotalSalesById cashier called", zap.Int32("id", req.GetCashierId()))

	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if month <= 0 || month > 12 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMonth
	}
	if id <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthTotalSalesCashier{
		Year:      year,
		Month:     month,
		CashierID: id,
	}

	methods, err := s.cashierStatsById.FindMonthlyTotalSalesById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthlyTotalSalesById cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthlyTotalSalesById cashier success")

	return &pb.ApiResponseCashierMonthlyTotalSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapResponseCashierMonthlyTotalSalesById(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindYearlyTotalSalesById(ctx context.Context, req *pb.FindYearTotalSalesById) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	s.logger.Info("FindYearlyTotalSalesById cashier called", zap.Int32("id", req.GetCashierId()))

	year := int(req.GetYear())
	id := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if id <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearTotalSalesCashier{
		Year:      year,
		CashierID: id,
	}

	methods, err := s.cashierStatsById.FindYearlyTotalSalesById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearlyTotalSalesById cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearlyTotalSalesById cashier success")

	return &pb.ApiResponseCashierYearlyTotalSales{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponseCashierYearlyTotalSalesById(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindMonthlyTotalSalesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalSalesByMerchant) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	s.logger.Info("FindMonthlyTotalSalesByMerchant cashier called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	month := int(req.GetMonth())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if month <= 0 || month > 12 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMonth
	}
	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.MonthTotalSalesMerchant{
		Year:       year,
		Month:      month,
		MerchantID: merchantId,
	}

	methods, err := s.cashierStatsByMerchant.FindMonthlyTotalSalesByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthlyTotalSalesByMerchant cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthlyTotalSalesByMerchant cashier success")

	return &pb.ApiResponseCashierMonthlyTotalSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapResponseCashierMonthlyTotalSalesByMerchant(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindYearlyTotalSalesByMerchant(ctx context.Context, req *pb.FindYearTotalSalesByMerchant) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	s.logger.Info("FindYearlyTotalSalesByMerchant cashier called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.YearTotalSalesMerchant{
		Year:       year,
		MerchantID: merchantId,
	}

	methods, err := s.cashierStatsByMerchant.FindYearlyTotalSalesByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearlyTotalSalesByMerchant cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearlyTotalSalesByMerchant cashier success")

	return &pb.ApiResponseCashierYearlyTotalSales{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponseCashierYearlyTotalSalesByMerchant(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindMonthSales(ctx context.Context, req *pb.FindYearCashier) (*pb.ApiResponseCashierMonthSales, error) {
	s.logger.Info("FindMonthSales cashier called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.cashierStats.FindMonthlySales(ctx, year)
	if err != nil {
		s.logger.Error("FindMonthSales cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthSales cashier success")

	return &pb.ApiResponseCashierMonthSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapResponsesCashierMonthlySales(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindYearSales(ctx context.Context, req *pb.FindYearCashier) (*pb.ApiResponseCashierYearSales, error) {
	s.logger.Info("FindYearSales cashier called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.cashierStats.FindYearlySales(ctx, year)
	if err != nil {
		s.logger.Error("FindYearSales cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearSales cashier success")

	return &pb.ApiResponseCashierYearSales{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponsesCashierYearlySales(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindMonthSalesByMerchant(ctx context.Context, req *pb.FindYearCashierByMerchant) (*pb.ApiResponseCashierMonthSales, error) {
	s.logger.Info("FindMonthSalesByMerchant cashier called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.MonthCashierMerchant{
		Year:       year,
		MerchantID: merchantId,
	}

	methods, err := s.cashierStatsByMerchant.FindMonthlyCashierByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthSalesByMerchant cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthSalesByMerchant cashier success")

	return &pb.ApiResponseCashierMonthSales{
		Status:  "success",
		Message: "Merchant monthly revenue retrieved successfully",
		Data:    mapResponsesCashierMonthlySalesByMerchant(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindYearSalesByMerchant(ctx context.Context, req *pb.FindYearCashierByMerchant) (*pb.ApiResponseCashierYearSales, error) {
	s.logger.Info("FindYearSalesByMerchant cashier called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.YearCashierMerchant{
		Year:       year,
		MerchantID: merchantId,
	}

	methods, err := s.cashierStatsByMerchant.FindYearlyCashierByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearSalesByMerchant cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearSalesByMerchant cashier success")

	return &pb.ApiResponseCashierYearSales{
		Status:  "success",
		Message: "Merchant yearly payment methods retrieved successfully",
		Data:    mapResponsesCashierYearlySalesByMerchant(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindMonthSalesById(ctx context.Context, req *pb.FindYearCashierById) (*pb.ApiResponseCashierMonthSales, error) {
	s.logger.Info("FindMonthSalesById cashier called", zap.Int32("id", req.GetCashierId()))

	year := int(req.GetYear())
	cashierId := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if cashierId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthCashierId{
		Year:      year,
		CashierID: cashierId,
	}

	methods, err := s.cashierStatsById.FindMonthlyCashierById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthSalesById cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthSalesById cashier success")

	return &pb.ApiResponseCashierMonthSales{
		Status:  "success",
		Message: "Cashier monthly sales retrieved successfully",
		Data:    mapResponsesCashierMonthlySalesById(methods),
	}, nil
}

func (s *cashierStatsHandleGrpc) FindYearSalesById(ctx context.Context, req *pb.FindYearCashierById) (*pb.ApiResponseCashierYearSales, error) {
	s.logger.Info("FindYearSalesById cashier called", zap.Int32("id", req.GetCashierId()))

	year := int(req.GetYear())
	cashierId := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}
	if cashierId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearCashierId{
		Year:      year,
		CashierID: cashierId,
	}

	methods, err := s.cashierStatsById.FindYearlyCashierById(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearSalesById cashier failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearSalesById cashier success")

	return &pb.ApiResponseCashierYearSales{
		Status:  "success",
		Message: "Cashier yearly sales retrieved successfully",
		Data:    mapResponsesCashierYearlySalesById(methods),
	}, nil
}
