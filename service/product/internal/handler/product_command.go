package handler

import (
	"context"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/product"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-product/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-shared/domain/requests"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/product_errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productCommandHandleGrpc struct {
	pb.UnimplementedProductCommandServiceServer
	productCommandService service.ProductCommandService
	logger                logger.LoggerInterface
}

func NewProductCommandHandleGrpc(
	service *service.Service,
	logger logger.LoggerInterface,
) pb.ProductCommandServiceServer {
	return &productCommandHandleGrpc{
		productCommandService: service.ProductCommand,
		logger:                logger,
	}
}

func (s *productCommandHandleGrpc) Create(ctx context.Context, request *pb.CreateProductRequest) (*pb.ApiResponseProduct, error) {
	s.logger.Info("Create product called", zap.String("name", request.GetName()))

	req := &requests.CreateProductRequest{
		MerchantID:   int(request.GetMerchantId()),
		CategoryID:   int(request.GetCategoryId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Price:        int(request.GetPrice()),
		CountInStock: int(request.GetCountInStock()),
		Brand:        request.GetBrand(),
		Weight:       int(request.GetWeight()),
		ImageProduct: request.GetImageProduct(),
	}

	if err := req.Validate(); err != nil {
		return nil, product_errors.ErrGrpcValidateCreateProduct
	}

	product, err := s.productCommandService.CreateProduct(ctx, req)
	if err != nil {
		s.logger.Error("Create product failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("Create product success")

	return &pb.ApiResponseProduct{
		Status:  "success",
		Message: "Successfully created product",
		Data:    mapResponseProduct(product),
	}, nil
}

func (s *productCommandHandleGrpc) Update(ctx context.Context, request *pb.UpdateProductRequest) (*pb.ApiResponseProduct, error) {
	s.logger.Info("Update product called", zap.Int32("id", request.GetProductId()))

	id := int(request.GetProductId())
	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	req := &requests.UpdateProductRequest{
		ProductID:    &id,
		MerchantID:   int(request.GetMerchantId()),
		CategoryID:   int(request.GetCategoryId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Price:        int(request.GetPrice()),
		CountInStock: int(request.GetCountInStock()),
		Brand:        request.GetBrand(),
		Weight:       int(request.GetWeight()),
		ImageProduct: request.GetImageProduct(),
	}

	if err := req.Validate(); err != nil {
		return nil, product_errors.ErrGrpcValidateUpdateProduct
	}

	product, err := s.productCommandService.UpdateProduct(ctx, req)
	if err != nil {
		s.logger.Error("Update product failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("Update product success")

	return &pb.ApiResponseProduct{
		Status:  "success",
		Message: "Successfully updated product",
		Data:    mapResponseProduct(product),
	}, nil
}

func (s *productCommandHandleGrpc) TrashedProduct(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDeleteAt, error) {
	s.logger.Info("TrashedProduct called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	product, err := s.productCommandService.TrashProduct(ctx, id)
	if err != nil {
		s.logger.Error("TrashedProduct failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("TrashedProduct success")

	return &pb.ApiResponseProductDeleteAt{
		Status:  "success",
		Message: "Successfully trashed product",
		Data:    mapResponseProductDeleteAt(product),
	}, nil
}

func (s *productCommandHandleGrpc) RestoreProduct(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDeleteAt, error) {
	s.logger.Info("RestoreProduct called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	product, err := s.productCommandService.RestoreProduct(ctx, id)
	if err != nil {
		s.logger.Error("RestoreProduct failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("RestoreProduct success")

	return &pb.ApiResponseProductDeleteAt{
		Status:  "success",
		Message: "Successfully restored product",
		Data:    mapResponseProductDeleteAt(product),
	}, nil
}

func (s *productCommandHandleGrpc) DeleteProductPermanent(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDelete, error) {
	s.logger.Info("DeleteProductPermanent called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	_, err := s.productCommandService.DeleteProductPermanent(ctx, id)
	if err != nil {
		s.logger.Error("DeleteProductPermanent failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("DeleteProductPermanent success")

	return &pb.ApiResponseProductDelete{
		Status:  "success",
		Message: "Successfully deleted Product permanently",
	}, nil
}

func (s *productCommandHandleGrpc) RestoreAllProduct(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseProductAll, error) {
	s.logger.Info("RestoreAllProduct called")

	_, err := s.productCommandService.RestoreAllProducts(ctx)
	if err != nil {
		s.logger.Error("RestoreAllProduct failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("RestoreAllProduct success")

	return &pb.ApiResponseProductAll{
		Status:  "success",
		Message: "Successfully restore all Product",
	}, nil
}

func (s *productCommandHandleGrpc) DeleteAllProductPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseProductAll, error) {
	s.logger.Info("DeleteAllProductPermanent called")

	_, err := s.productCommandService.DeleteAllProductsPermanent(ctx)
	if err != nil {
		s.logger.Error("DeleteAllProductPermanent failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("DeleteAllProductPermanent success")

	return &pb.ApiResponseProductAll{
		Status:  "success",
		Message: "Successfully delete Product permanen",
	}, nil
}
