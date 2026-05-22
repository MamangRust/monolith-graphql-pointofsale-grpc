package handler

import (
	"context"

	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"
	"github.com/MamangRust/monolith-point-of-sale-shared/domain/requests"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors"
	"github.com/MamangRust/monolith-point-of-sale-shared/errors/user_errors"
	"github.com/MamangRust/monolith-point-of-sale-user/internal/service"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
)

type userCommandHandleGrpc struct {
	pb.UnimplementedUserCommandServiceServer
	userCommandService service.UserCommandService
	logger             logger.LoggerInterface
}

func NewUserCommandHandleGrpc(
	user *service.Service,
	logger logger.LoggerInterface,
) pb.UserCommandServiceServer {
	return &userCommandHandleGrpc{
		userCommandService: user.UserCommand,
		logger:             logger,
	}
}

func (s *userCommandHandleGrpc) Create(ctx context.Context, request *pb.CreateUserRequest) (*pb.ApiResponseUser, error) {
	s.logger.Info("Create user called", zap.String("email", request.GetEmail()))

	req := &requests.CreateUserRequest{
		FirstName:       request.GetFirstname(),
		LastName:        request.GetLastname(),
		Email:           request.GetEmail(),
		Password:        request.GetPassword(),
		ConfirmPassword: request.GetConfirmPassword(),
	}

	if err := req.Validate(); err != nil {
		return nil, user_errors.ErrGrpcValidateCreateUser
	}

	user, err := s.userCommandService.CreateUser(ctx, req)
	if err != nil {
		s.logger.Error("Create user failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("Create user success")

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully created user",
		Data:    mapUserToProto(user),
	}, nil
}

func (s *userCommandHandleGrpc) Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.ApiResponseUser, error) {
	s.logger.Info("Update user called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, user_errors.ErrGrpcUserInvalidId
	}

	req := &requests.UpdateUserRequest{
		UserID:          &id,
		FirstName:       request.GetFirstname(),
		LastName:        request.GetLastname(),
		Email:           request.GetEmail(),
		Password:        request.GetPassword(),
		ConfirmPassword: request.GetConfirmPassword(),
	}

	if err := req.Validate(); err != nil {
		return nil, user_errors.ErrGrpcValidateCreateUser
	}

	user, err := s.userCommandService.UpdateUser(ctx, req)
	if err != nil {
		s.logger.Error("Update user failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("Update user success")

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully updated user",
		Data:    mapUserToProto(user),
	}, nil
}

func (s *userCommandHandleGrpc) TrashedUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDeleteAt, error) {
	s.logger.Info("TrashedUser called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, user_errors.ErrGrpcUserInvalidId
	}

	user, err := s.userCommandService.TrashedUser(ctx, id)
	if err != nil {
		s.logger.Error("TrashedUser failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("TrashedUser success")

	return &pb.ApiResponseUserDeleteAt{
		Status:  "success",
		Message: "Successfully trashed user",
		Data:    mapUserDeleteAtToProto(user),
	}, nil
}

func (s *userCommandHandleGrpc) RestoreUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDeleteAt, error) {
	s.logger.Info("RestoreUser called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, user_errors.ErrGrpcUserInvalidId
	}

	user, err := s.userCommandService.RestoreUser(ctx, id)
	if err != nil {
		s.logger.Error("RestoreUser failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("RestoreUser success")

	return &pb.ApiResponseUserDeleteAt{
		Status:  "success",
		Message: "Successfully restored user",
		Data:    mapUserDeleteAtToProto(user),
	}, nil
}

func (s *userCommandHandleGrpc) DeleteUserPermanent(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDelete, error) {
	s.logger.Info("DeleteUserPermanent called", zap.Int32("id", request.GetId()))

	id := int(request.GetId())
	if id == 0 {
		return nil, user_errors.ErrGrpcUserInvalidId
	}

	_, err := s.userCommandService.DeleteUserPermanent(ctx, id)
	if err != nil {
		s.logger.Error("DeleteUserPermanent failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("DeleteUserPermanent success")

	return &pb.ApiResponseUserDelete{
		Status:  "success",
		Message: "Successfully deleted user permanently",
	}, nil
}

func (s *userCommandHandleGrpc) RestoreAllUser(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseUserAll, error) {
	s.logger.Info("RestoreAllUser called")

	_, err := s.userCommandService.RestoreAllUser(ctx)
	if err != nil {
		s.logger.Error("RestoreAllUser failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("RestoreAllUser success")

	return &pb.ApiResponseUserAll{
		Status:  "success",
		Message: "Successfully restore all user",
	}, nil
}

func (s *userCommandHandleGrpc) DeleteAllUserPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseUserAll, error) {
	s.logger.Info("DeleteAllUserPermanent called")

	_, err := s.userCommandService.DeleteAllUserPermanent(ctx)
	if err != nil {
		s.logger.Error("DeleteAllUserPermanent failed", zap.Error(err))
		return nil, errors.ToGrpcError(err)
	}

	s.logger.Info("DeleteAllUserPermanent success")

	return &pb.ApiResponseUserAll{
		Status:  "success",
		Message: "Successfully delete user permanen",
	}, nil
}
