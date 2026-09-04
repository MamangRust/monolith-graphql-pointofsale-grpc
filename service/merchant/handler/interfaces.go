package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pbmerchant_document "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant_document"
)

type MerchantDocumentHandleGrpc interface {
	pbmerchant_document.MerchantDocumentQueryServiceServer
	pbmerchant_document.MerchantDocumentCommandServiceServer
}

type MerchantHandleGrpc interface {
	pb.MerchantQueryServiceServer
	pb.MerchantCommandServiceServer
}
