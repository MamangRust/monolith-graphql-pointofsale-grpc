package handler

import (
	"context"
	"math"

	apipb "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/category"
	"github.com/MamangRust/monolith-point-of-sale-category/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-shared/domain/requests"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/category_errors"
	"go.uber.org/zap"
)

type categoryQueryHandleGrpc struct {
	pb.UnimplementedCategoryQueryServiceServer
	categoryQueryService service.CategoryQueryService
	logger               logger.LoggerInterface
}

func NewcategoryQueryHandleGrpc(categoryQueryService service.CategoryQueryService, logger logger.LoggerInterface) pb.CategoryQueryServiceServer {
	return &categoryQueryHandleGrpc{
		categoryQueryService: categoryQueryService,
		logger:               logger,
	}
}

func (s *categoryQueryHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategory, error) {
	s.logger.Info("FindAll categories called", zap.Int32("page", request.GetPage()))

	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCategory{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	category, totalRecords, err := s.categoryQueryService.FindAll(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindAll categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &apipb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	s.logger.Info("FindAll categories success", zap.Int("count", len(category)))

	return &pb.ApiResponsePaginationCategory{
		Status:     "success",
		Message:    "Successfully fetched categories",
		Data:       mapResponsesCategory(category),
		Pagination: mapPaginationMeta(paginationMeta),
	}, nil
}

func (s *categoryQueryHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategory, error) {
	s.logger.Info("FindById category called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryQueryService.FindById(ctx, id)
	if err != nil {
		s.logger.Error("FindById category failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindById category success", zap.Int("id", id))

	return &pb.ApiResponseCategory{
		Status:  "success",
		Message: "Successfully fetched category",
		Data:    mapResponseCategory(category),
	}, nil
}

func (s *categoryQueryHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategoryDeleteAt, error) {
	s.logger.Info("FindByActive categories called", zap.Int32("page", request.GetPage()))

	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCategory{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	categories, totalRecords, err := s.categoryQueryService.FindByActive(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindByActive categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &apipb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	s.logger.Info("FindByActive categories success")

	return &pb.ApiResponsePaginationCategoryDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active categories",
		Data:       mapResponsesCategoryActive(categories),
		Pagination: mapPaginationMeta(paginationMeta),
	}, nil
}

func (s *categoryQueryHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategoryDeleteAt, error) {
	s.logger.Info("FindByTrashed categories called")

	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCategory{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	categories, totalRecords, err := s.categoryQueryService.FindByTrashed(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindByTrashed categories failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &apipb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	s.logger.Info("FindByTrashed categories success")

	return &pb.ApiResponsePaginationCategoryDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed categories",
		Data:       mapResponsesCategoryTrashed(categories),
		Pagination: mapPaginationMeta(paginationMeta),
	}, nil
}
