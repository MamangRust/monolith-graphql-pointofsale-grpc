package repository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	sharedErrors "github.com/MamangRust/monolith-graphql-pointofsale-shared/errors"
	"github.com/jackc/pgx/v5/pgtype"

	pbuser "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
)

type userQueryRepository struct {
	client pbuser.UserQueryServiceClient
}

func NewUserQueryRepository(client pbuser.UserQueryServiceClient) UserQueryRepository {
	return &userQueryRepository{
		client: client,
	}
}

func (r *userQueryRepository) FindById(ctx context.Context, userID int) (*db.User, error) {
	res, err := r.client.FindById(ctx, &pbuser.FindByIdUserRequest{
		Id: int32(userID),
	})
	if err != nil || res == nil || res.Data == nil {
		return nil, sharedErrors.ErrInternal
	}

	var createdAt, updatedAt pgtype.Timestamp
	if res.Data.CreatedAt != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", res.Data.CreatedAt); err == nil {
			createdAt = pgtype.Timestamp{Time: t, Valid: true}
		} else if t, err = time.Parse(time.RFC3339, res.Data.CreatedAt); err == nil {
			createdAt = pgtype.Timestamp{Time: t, Valid: true}
		}
	}
	if res.Data.UpdatedAt != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", res.Data.UpdatedAt); err == nil {
			updatedAt = pgtype.Timestamp{Time: t, Valid: true}
		} else if t, err = time.Parse(time.RFC3339, res.Data.UpdatedAt); err == nil {
			updatedAt = pgtype.Timestamp{Time: t, Valid: true}
		}
	}

	return &db.User{
		UserID:    res.Data.Id,
		Firstname: res.Data.Firstname,
		Lastname:  res.Data.Lastname,
		Email:     res.Data.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
