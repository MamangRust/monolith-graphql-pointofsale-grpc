package handler

import (
	pbutils "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
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

func mapUserToProto(user *db.User) *pb.UserResponse {
	if user == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	if user.CreatedAt.Valid {
		createdAtStr = user.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.UpdatedAt.Valid {
		updatedAtStr = user.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &pb.UserResponse{
		Id:        int32(user.UserID),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
	}
}

func mapGetUsersRowToProto(user *db.GetUsersRow) *pb.UserResponse {
	if user == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	if user.CreatedAt.Valid {
		createdAtStr = user.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.UpdatedAt.Valid {
		updatedAtStr = user.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &pb.UserResponse{
		Id:        int32(user.UserID),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
	}
}

func mapGetUsersRowsToProto(users []*db.GetUsersRow) []*pb.UserResponse {
	var responseUsers []*pb.UserResponse
	for _, u := range users {
		responseUsers = append(responseUsers, mapGetUsersRowToProto(u))
	}
	return responseUsers
}

func mapUserDeleteAtToProto(user *db.User) *pb.UserResponseDeleteAt {
	if user == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	var deletedAt *wrapperspb.StringValue
	if user.CreatedAt.Valid {
		createdAtStr = user.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.UpdatedAt.Valid {
		updatedAtStr = user.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.DeletedAt.Valid {
		deletedAt = wrapperspb.String(user.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}
	return &pb.UserResponseDeleteAt{
		Id:        int32(user.UserID),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
		DeletedAt: deletedAt,
	}
}

func mapActiveUserToProto(user *db.GetUsersActiveRow) *pb.UserResponseDeleteAt {
	if user == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	var deletedAt *wrapperspb.StringValue
	if user.CreatedAt.Valid {
		createdAtStr = user.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.UpdatedAt.Valid {
		updatedAtStr = user.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.DeletedAt.Valid {
		deletedAt = wrapperspb.String(user.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}
	return &pb.UserResponseDeleteAt{
		Id:        int32(user.UserID),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
		DeletedAt: deletedAt,
	}
}

func mapActiveUsersToProto(users []*db.GetUsersActiveRow) []*pb.UserResponseDeleteAt {
	var responseUsers []*pb.UserResponseDeleteAt
	for _, u := range users {
		responseUsers = append(responseUsers, mapActiveUserToProto(u))
	}
	return responseUsers
}

func mapTrashedUserToProto(user *db.GetUserTrashedRow) *pb.UserResponseDeleteAt {
	if user == nil {
		return nil
	}
	var createdAtStr, updatedAtStr string
	var deletedAt *wrapperspb.StringValue
	if user.CreatedAt.Valid {
		createdAtStr = user.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.UpdatedAt.Valid {
		updatedAtStr = user.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.DeletedAt.Valid {
		deletedAt = wrapperspb.String(user.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}
	return &pb.UserResponseDeleteAt{
		Id:        int32(user.UserID),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
		DeletedAt: deletedAt,
	}
}

func mapTrashedUsersToProto(users []*db.GetUserTrashedRow) []*pb.UserResponseDeleteAt {
	var responseUsers []*pb.UserResponseDeleteAt
	for _, u := range users {
		responseUsers = append(responseUsers, mapTrashedUserToProto(u))
	}
	return responseUsers
}
