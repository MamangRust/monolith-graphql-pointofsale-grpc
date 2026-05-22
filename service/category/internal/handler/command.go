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
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryCommandHandleGrpc struct {
	pb.UnimplementedCategoryCommandServiceServer
	categoryCommandService service.CategoryCommandService
	logger                 logger.LoggerInterface
}

func NewcategoryCommandHandleGrpc(categoryCommandService service.CategoryCommandService, logger logger.LoggerInterface) pb.CategoryCommandServiceServer {
	return &categoryCommandHandleGrpc{
		categoryCommandService: categoryCommandService,
		logger:                 logger,
	}
}

func (s *categoryCommandHandleGrpc) Create(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.ApiResponseCategory, error) {
	s.logger.Info("Create category called", zap.String("name", request.GetName()))

	req := &requests.CreateCategoryRequest{
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	if err := req.Validate(); err != nil {
		return nil, category_errors.ErrGrpcValidateCreateCategory
	}

	category, err := s.categoryCommandService.CreateCategory(ctx, req)
	if err != nil {
		s.logger.Error("Create category failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("Create category success")

	return &pb.ApiResponseCategory{
		Status:  "success",
		Message: "Successfully created category",
		Data:    mapResponseCategory(category),
	}, nil
}

func (s *categoryCommandHandleGrpc) Update(ctx context.Context, request *pb.UpdateCategoryRequest) (*pb.ApiResponseCategory, error) {
	s.logger.Info("Update category called", zap.Int32("id", request.GetCategoryId()))

	id := int(request.GetCategoryId())
	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	req := &requests.UpdateCategoryRequest{
		CategoryID:  &id,
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	if err := req.Validate(); err != nil {
		return nil, category_errors.ErrGrpcValidateUpdateCategory
	}

	category, err := s.categoryCommandService.UpdateCategory(ctx, req)
	if err != nil {
		s.logger.Error("Update category failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("Update category success")

	return &pb.ApiResponseCategory{
		Status:  "success",
		Message: "Successfully updated category",
		Data:    mapResponseCategory(category),
	}, nil
}

func (s *categoryCommandHandleGrpc) TrashedCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	s.logger.Info("TrashedCategory called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryCommandService.TrashedCategory(ctx, id)
	if err != nil {
		s.logger.Error("TrashedCategory failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("TrashedCategory success")

	return &pb.ApiResponseCategoryDeleteAt{
		Status:  "success",
		Message: "Successfully trashed category",
		Data:    mapResponseCategoryDeleteAt(category),
	}, nil
}

func (s *categoryCommandHandleGrpc) RestoreCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	s.logger.Info("RestoreCategory called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryCommandService.RestoreCategory(ctx, id)
	if err != nil {
		s.logger.Error("RestoreCategory failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("RestoreCategory success")

	return &pb.ApiResponseCategoryDeleteAt{
		Status:  "success",
		Message: "Successfully restored category",
		Data:    mapResponseCategoryDeleteAt(category),
	}, nil
}

func (s *categoryCommandHandleGrpc) DeleteCategoryPermanent(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDelete, error) {
	s.logger.Info("DeleteCategoryPermanent called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	_, err := s.categoryCommandService.DeleteCategoryPermanent(ctx, id)
	if err != nil {
		s.logger.Error("DeleteCategoryPermanent failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("DeleteCategoryPermanent success")

	return &pb.ApiResponseCategoryDelete{
		Status:  "success",
		Message: "Successfully deleted category permanently",
	}, nil
}

func (s *categoryCommandHandleGrpc) RestoreAllCategory(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	s.logger.Info("RestoreAllCategory called")

	_, err := s.categoryCommandService.RestoreAllCategories(ctx)
	if err != nil {
		s.logger.Error("RestoreAllCategory failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("RestoreAllCategory success")

	return &pb.ApiResponseCategoryAll{
		Status:  "success",
		Message: "Successfully restore all category",
	}, nil
}

func (s *categoryCommandHandleGrpc) DeleteAllCategoryPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	s.logger.Info("DeleteAllCategoryPermanent called")

	_, err := s.categoryCommandService.DeleteAllCategoriesPermanent(ctx)
	if err != nil {
		s.logger.Error("DeleteAllCategoryPermanent failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("DeleteAllCategoryPermanent success")

	return &pb.ApiResponseCategoryAll{
		Status:  "success",
		Message: "Successfully delete category permanen",
	}, nil
}
