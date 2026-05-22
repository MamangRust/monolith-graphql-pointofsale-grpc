package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	"github.com/MamangRust/monolith-point-of-sale-merchant/internal/service"
	"github.com/MamangRust/monolith-point-of-sale-pkg/logger"

	pbdocument "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant_document"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	MerchantQuery           pb.MerchantQueryServiceServer
	MerchantCommand         pb.MerchantCommandServiceServer
	MerchantDocumentCommand pbdocument.MerchantDocumentCommandServiceServer
	MerchantDocumentQuery   pbdocument.MerchantDocumentQueryServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		MerchantQuery: NewMerchantQueryGrpc(
			deps.Service,
			deps.Logger,
		),
		MerchantCommand: NewMerchantCommandGrpc(
			deps.Service,
			deps.Logger,
		),
		MerchantDocumentQuery: NewMerchantDocumentQueryGrpc(
			deps.Service,
			deps.Logger,
		),
		MerchantDocumentCommand: NewMerchantDocumentCommandGrpc(
			deps.Service,
			deps.Logger,
		),
	}
}
