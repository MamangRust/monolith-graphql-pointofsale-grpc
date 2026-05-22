package repository

import (
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/user"
)

type Repositories struct {
	MerchantQuery           MerchantQueryRepository
	MerchantCommand         MerchantCommandRepository
	MerchantDocumentCommand MerchantDocumentCommandRepository
	MerchantDocumentQuery   MerchantDocumentQueryRepository
	UserQuery               UserQueryRepository
}

func NewRepositories(
	DB *db.Queries,
	userClient pb.UserQueryServiceClient,
) *Repositories {
	return &Repositories{
		MerchantQuery:           NewMerchantQueryRepository(DB),
		MerchantCommand:         NewMerchantCommandRepository(DB),
		MerchantDocumentCommand: NewMerchantDocumentCommandRepository(DB),
		MerchantDocumentQuery:   NewMerchantDocumentQueryRepository(DB),
		UserQuery:               NewUserQueryRepository(userClient),
	}
}
