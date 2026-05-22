package handler

import (
	"context"
	"math"

	pbutils "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-shared/domain/requests"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/user_errors"
	"github.com/MamangRust/monolith-point-of-sale-user/internal/service"
	"go.uber.org/zap"
)

type userQueryHandleGrpc struct {
	pb.UnimplementedUserQueryServiceServer
	userQueryService service.UserQueryService
	logger           logger.LoggerInterface
}

func NewUserQueryHandleGrpc(
	user *service.Service,
	logger logger.LoggerInterface,
) pb.UserQueryServiceServer {
	return &userQueryHandleGrpc{
		userQueryService: user.UserQuery,
		logger:           logger,
	}
}

func (s *userQueryHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUser, error) {
	s.logger.Info("FindAll users called", zap.Int32("page", request.GetPage()))

	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllUsers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	users, totalRecords, err := s.userQueryService.FindAll(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindAll users failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	var totalRecordsVal int
	if totalRecords != nil {
		totalRecordsVal = *totalRecords
	}
	totalPages := int(math.Ceil(float64(totalRecordsVal) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecordsVal),
	}

	s.logger.Info("FindAll users success")

	return &pb.ApiResponsePaginationUser{
		Status:     "success",
		Message:    "Successfully fetched users",
		Data:       mapGetUsersRowsToProto(users),
		Pagination: mapPaginationMeta(paginationMeta),
	}, nil
}

func (s *userQueryHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error) {
	s.logger.Info("FindById user called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, user_errors.ErrGrpcUserNotFound
	}

	user, err := s.userQueryService.FindByID(ctx, id)
	if err != nil {
		s.logger.Error("FindById user failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("FindById user success")

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully fetched user",
		Data:    mapUserToProto(user),
	}, nil
}

func (s *userQueryHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	s.logger.Info("FindByActive users called", zap.Int32("page", request.GetPage()))

	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllUsers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	users, totalRecords, err := s.userQueryService.FindByActive(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindByActive users failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	var totalRecordsVal int
	if totalRecords != nil {
		totalRecordsVal = *totalRecords
	}
	totalPages := int(math.Ceil(float64(totalRecordsVal) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecordsVal),
	}

	s.logger.Info("FindByActive users success")

	return &pb.ApiResponsePaginationUserDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active users",
		Data:       mapActiveUsersToProto(users),
		Pagination: mapPaginationMeta(paginationMeta),
	}, nil
}

func (s *userQueryHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	s.logger.Info("FindByTrashed users called")

	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllUsers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	users, totalRecords, err := s.userQueryService.FindByTrashed(ctx, &reqService)
	if err != nil {
		s.logger.Error("FindByTrashed users failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	var totalRecordsVal int
	if totalRecords != nil {
		totalRecordsVal = *totalRecords
	}
	totalPages := int(math.Ceil(float64(totalRecordsVal) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecordsVal),
	}

	s.logger.Info("FindByTrashed users success")

	return &pb.ApiResponsePaginationUserDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed users",
		Data:       mapTrashedUsersToProto(users),
		Pagination: mapPaginationMeta(paginationMeta),
	}, nil
}
