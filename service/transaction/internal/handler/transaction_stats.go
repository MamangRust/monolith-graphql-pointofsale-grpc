package handler

import (
	"context"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/transaction"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-shared/domain/requests"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/transaction_errors"
	"github.com/MamangRust/monolith-point-of-sale-transacton/internal/service"
	"go.uber.org/zap"
)

type transactionStatsHandleGrpc struct {
	pb.UnimplementedTransactionStatsServiceServer
	transactionStats           service.TransactionStatsService
	transactionStatsByMerchant service.TransactionStatsByMerchantService
	logger                     logger.LoggerInterface
}

func NewTransactionStatsHandleGrpc(transactionStatsService service.TransactionStatsService, transactionStatsByMerchant service.TransactionStatsByMerchantService, logger logger.LoggerInterface) *transactionStatsHandleGrpc {
	return &transactionStatsHandleGrpc{
		transactionStats:           transactionStatsService,
		transactionStatsByMerchant: transactionStatsByMerchant,
		logger:                     logger,
	}
}

func (s *transactionStatsHandleGrpc) FindMonthStatusSuccess(ctx context.Context, request *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	s.logger.Info("FindMonthStatusSuccess transactions called", zap.Int32("year", request.GetYear()))

	year := int(request.GetYear())
	month := int(request.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthAmountTransaction{
		Year:  year,
		Month: month,
	}

	res, err := s.transactionStats.FindMonthlyAmountSuccess(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthStatusSuccess transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthStatusSuccess transactions success")

	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  "success",
		Message: "Monthly success data retrieved successfully",
		Data:    mapResponsesTransactionMonthlyAmountSuccess(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearStatusSuccess(ctx context.Context, request *pb.FindYearlyTransactionStatus) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	s.logger.Info("FindYearStatusSuccess transactions called", zap.Int32("year", request.GetYear()))

	year := int(request.GetYear())
	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	res, err := s.transactionStats.FindYearlyAmountSuccess(ctx, year)
	if err != nil {
		s.logger.Error("FindYearStatusSuccess transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearStatusSuccess transactions success")

	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  "success",
		Message: "Yearly success data retrieved successfully",
		Data:    mapResponsesTransactionYearlyAmountSuccess(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindMonthStatusFailed(ctx context.Context, request *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	s.logger.Info("FindMonthStatusFailed transactions called", zap.Int32("year", request.GetYear()))

	year := int(request.GetYear())
	month := int(request.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthAmountTransaction{
		Year:  year,
		Month: month,
	}

	res, err := s.transactionStats.FindMonthlyAmountFailed(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthStatusFailed transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthStatusFailed transactions success")

	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  "success",
		Message: "Monthly failed data retrieved successfully",
		Data:    mapResponsesTransactionMonthlyAmountFailed(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearStatusFailed(ctx context.Context, request *pb.FindYearlyTransactionStatus) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	s.logger.Info("FindYearStatusFailed transactions called", zap.Int32("year", request.GetYear()))

	year := int(request.GetYear())
	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	res, err := s.transactionStats.FindYearlyAmountFailed(ctx, year)
	if err != nil {
		s.logger.Error("FindYearStatusFailed transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearStatusFailed transactions success")

	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  "success",
		Message: "Yearly failed data retrieved successfully",
		Data:    mapResponsesTransactionYearlyAmountFailed(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindMonthStatusSuccessByMerchant(ctx context.Context, request *pb.FindMonthlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	s.logger.Info("FindMonthStatusSuccessByMerchant transactions called", zap.Int32("merchantId", request.GetMerchantId()))

	year := int(request.GetYear())
	month := int(request.GetMonth())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.MonthAmountTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	res, err := s.transactionStatsByMerchant.FindMonthlyAmountSuccessByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthStatusSuccessByMerchant transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthStatusSuccessByMerchant transactions success")

	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  "success",
		Message: "Merchant monthly success data retrieved successfully",
		Data:    mapResponsesTransactionMonthlyAmountSuccessByMerchant(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearStatusSuccessByMerchant(ctx context.Context, request *pb.FindYearlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	s.logger.Info("FindYearStatusSuccessByMerchant transactions called", zap.Int32("merchantId", request.GetMerchantId()))

	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearAmountTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	res, err := s.transactionStatsByMerchant.FindYearlyAmountSuccessByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearStatusSuccessByMerchant transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearStatusSuccessByMerchant transactions success")

	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  "success",
		Message: "Merchant yearly success data retrieved successfully",
		Data:    mapResponsesTransactionYearlyAmountSuccessByMerchant(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindMonthStatusFailedByMerchant(ctx context.Context, request *pb.FindMonthlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	s.logger.Info("FindMonthStatusFailedByMerchant transactions called", zap.Int32("merchantId", request.GetMerchantId()))

	year := int(request.GetYear())
	month := int(request.GetMonth())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.MonthAmountTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	res, err := s.transactionStatsByMerchant.FindMonthlyAmountFailedByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthStatusFailedByMerchant transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthStatusFailedByMerchant transactions success")

	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  "success",
		Message: "Merchant monthly failed data retrieved successfully",
		Data:    mapResponsesTransactionMonthlyAmountFailedByMerchant(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearStatusFailedByMerchant(ctx context.Context, request *pb.FindYearlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	s.logger.Info("FindYearStatusFailedByMerchant transactions called", zap.Int32("merchantId", request.GetMerchantId()))

	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearAmountTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	res, err := s.transactionStatsByMerchant.FindYearlyAmountFailedByMerchant(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearStatusFailedByMerchant transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearStatusFailedByMerchant transactions success")

	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  "success",
		Message: "Merchant yearly failed data retrieved successfully",
		Data:    mapResponsesTransactionYearlyAmountFailedByMerchant(res),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindMonthMethodSuccess(ctx context.Context, req *pb.MonthTransactionMethod) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	s.logger.Info("FindMonthMethodSuccess transactions called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	methods, err := s.transactionStats.FindMonthlyMethodSuccess(ctx, &requests.MonthMethodTransaction{
		Year:  year,
		Month: month,
	})
	if err != nil {
		s.logger.Error("FindMonthMethodSuccess transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthMethodSuccess transactions success")

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "Monthly payment methods retrieved successfully",
		Data:    mapResponsesTransactionMonthlyMethodSuccess(methods),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearMethodSuccess(ctx context.Context, req *pb.YearTransactionMethod) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	s.logger.Info("FindYearMethodSuccess transactions called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	methods, err := s.transactionStats.FindYearlyMethodSuccess(ctx, year)
	if err != nil {
		s.logger.Error("FindYearMethodSuccess transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearMethodSuccess transactions success")

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapResponsesTransactionYearlyMethodSuccess(methods),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindMonthMethodByMerchantSuccess(ctx context.Context, req *pb.MonthTransactionMethodByMerchant) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	s.logger.Info("FindMonthMethodByMerchantSuccess transactions called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	id := int(req.GetMerchantId())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthMethodTransactionMerchant{
		Year:       year,
		MerchantID: id,
		Month:      month,
	}

	methods, err := s.transactionStatsByMerchant.FindMonthlyMethodByMerchantSuccess(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthMethodByMerchantSuccess transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthMethodByMerchantSuccess transactions success")

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "Merchant monthly payment methods retrieved successfully",
		Data:    mapResponsesTransactionMonthlyMethodByMerchantSuccess(methods),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearMethodByMerchantSuccess(ctx context.Context, req *pb.YearTransactionMethodByMerchant) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	s.logger.Info("FindYearMethodByMerchantSuccess transactions called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearMethodTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.transactionStatsByMerchant.FindYearlyMethodByMerchantSuccess(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearMethodByMerchantSuccess transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearMethodByMerchantSuccess transactions success")

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "Merchant yearly payment methods retrieved successfully",
		Data:    mapResponsesTransactionYearlyMethodByMerchantSuccess(methods),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindMonthMethodFailed(ctx context.Context, req *pb.MonthTransactionMethod) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	s.logger.Info("FindMonthMethodFailed transactions called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	methods, err := s.transactionStats.FindMonthlyMethodFailed(ctx, &requests.MonthMethodTransaction{
		Year:  year,
		Month: month,
	})
	if err != nil {
		s.logger.Error("FindMonthMethodFailed transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthMethodFailed transactions success")

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "Failed",
		Message: "Monthly payment methods retrieved Failedfully",
		Data:    mapResponsesTransactionMonthlyMethodFailed(methods),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearMethodFailed(ctx context.Context, req *pb.YearTransactionMethod) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	s.logger.Info("FindYearMethodFailed transactions called", zap.Int32("year", req.GetYear()))

	year := int(req.GetYear())
	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	methods, err := s.transactionStats.FindYearlyMethodFailed(ctx, year)
	if err != nil {
		s.logger.Error("FindYearMethodFailed transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearMethodFailed transactions success")

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "Failed",
		Message: "Yearly payment methods retrieved Failedfully",
		Data:    mapResponsesTransactionYearlyMethodFailed(methods),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindMonthMethodByMerchantFailed(ctx context.Context, req *pb.MonthTransactionMethodByMerchant) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	s.logger.Info("FindMonthMethodByMerchantFailed transactions called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	id := int(req.GetMerchantId())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}
	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthMethodTransactionMerchant{
		Year:       year,
		MerchantID: id,
		Month:      month,
	}

	methods, err := s.transactionStatsByMerchant.FindMonthlyMethodByMerchantFailed(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindMonthMethodByMerchantFailed transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindMonthMethodByMerchantFailed transactions success")

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "Failed",
		Message: "Merchant monthly payment methods retrieved Failedfully",
		Data:    mapResponsesTransactionMonthlyMethodByMerchantFailed(methods),
	}, nil
}

func (s *transactionStatsHandleGrpc) FindYearlyMethodByMerchantFailed(ctx context.Context, req *pb.YearTransactionMethodByMerchant) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	s.logger.Info("FindYearlyMethodByMerchantFailed transactions called", zap.Int32("merchantId", req.GetMerchantId()))

	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearMethodTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.transactionStatsByMerchant.FindYearlyMethodByMerchantFailed(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindYearlyMethodByMerchantFailed transactions failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindYearlyMethodByMerchantFailed transactions success")

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "Failed",
		Message: "Merchant yearly payment methods retrieved Failedfully",
		Data:    mapResponsesTransactionYearlyMethodByMerchantFailed(methods),
	}, nil
}
