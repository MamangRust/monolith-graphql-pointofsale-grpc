package handler

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-merchant/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
	pbmerchant "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pbmerchant_document "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant_document"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Merchant                pbmerchant.MerchantQueryServiceServer
	MerchantCommand         pbmerchant.MerchantCommandServiceServer
	MerchantDocument        pbmerchant_document.MerchantDocumentQueryServiceServer
	MerchantDocumentCommand pbmerchant_document.MerchantDocumentCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	h := NewMerchantHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	hd := NewMerchantDocumentHandleGrpc(
		deps.Service,
		deps.Logger,
	)
	return &Handler{
		Merchant:                h,
		MerchantCommand:         h,
		MerchantDocument:        hd,
		MerchantDocumentCommand: hd,
	}
}
