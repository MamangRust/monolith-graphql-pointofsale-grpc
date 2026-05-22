package handler

import (
	apipb "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/role"
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
)

func mapPaginationMeta(meta *apipb.PaginationMeta) *apipb.PaginationMeta {
	if meta == nil {
		return nil
	}
	return &apipb.PaginationMeta{
		CurrentPage:  meta.CurrentPage,
		PageSize:     meta.PageSize,
		TotalPages:   meta.TotalPages,
		TotalRecords: meta.TotalRecords,
	}
}

func mapResponseRole(role *db.Role) *pb.RoleResponse {
	if role == nil {
		return nil
	}
	var createdAtStr string
	if role.CreatedAt.Valid {
		createdAtStr = role.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var updatedAtStr string
	if role.UpdatedAt.Valid {
		updatedAtStr = role.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &pb.RoleResponse{
		Id:        role.RoleID,
		Name:      role.RoleName,
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
	}
}

func mapResponsesRole(roles []*db.GetRolesRow) []*pb.RoleResponse {
	var responseRoles []*pb.RoleResponse
	for _, role := range roles {
		if role == nil {
			continue
		}
		var createdAtStr string
		if role.CreatedAt.Valid {
			createdAtStr = role.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var updatedAtStr string
		if role.UpdatedAt.Valid {
			updatedAtStr = role.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		responseRoles = append(responseRoles, &pb.RoleResponse{
			Id:        role.RoleID,
			Name:      role.RoleName,
			CreatedAt: createdAtStr,
			UpdatedAt: updatedAtStr,
		})
	}
	return responseRoles
}

func mapResponsesRoleFromDB(roles []*db.Role) []*pb.RoleResponse {
	var responseRoles []*pb.RoleResponse
	for _, role := range roles {
		responseRoles = append(responseRoles, mapResponseRole(role))
	}
	return responseRoles
}

func mapResponsesRoleFromActive(roles []*db.GetActiveRolesRow) []*pb.RoleResponseDeleteAt {
	var responseRoles []*pb.RoleResponseDeleteAt
	for _, role := range roles {
		if role == nil {
			continue
		}
		var createdAtStr string
		if role.CreatedAt.Valid {
			createdAtStr = role.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var updatedAtStr string
		if role.UpdatedAt.Valid {
			updatedAtStr = role.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAtStr string
		if role.DeletedAt.Valid {
			deletedAtStr = role.DeletedAt.Time.Format("2006-01-02 15:04:05")
		}
		responseRoles = append(responseRoles, &pb.RoleResponseDeleteAt{
			Id:        role.RoleID,
			Name:      role.RoleName,
			CreatedAt: createdAtStr,
			UpdatedAt: updatedAtStr,
			DeletedAt: deletedAtStr,
		})
	}
	return responseRoles
}

func mapResponsesRoleFromTrashed(roles []*db.GetTrashedRolesRow) []*pb.RoleResponseDeleteAt {
	var responseRoles []*pb.RoleResponseDeleteAt
	for _, role := range roles {
		if role == nil {
			continue
		}
		var createdAtStr string
		if role.CreatedAt.Valid {
			createdAtStr = role.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var updatedAtStr string
		if role.UpdatedAt.Valid {
			updatedAtStr = role.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAtStr string
		if role.DeletedAt.Valid {
			deletedAtStr = role.DeletedAt.Time.Format("2006-01-02 15:04:05")
		}
		responseRoles = append(responseRoles, &pb.RoleResponseDeleteAt{
			Id:        role.RoleID,
			Name:      role.RoleName,
			CreatedAt: createdAtStr,
			UpdatedAt: updatedAtStr,
			DeletedAt: deletedAtStr,
		})
	}
	return responseRoles
}
