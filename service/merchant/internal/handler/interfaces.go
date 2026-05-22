package handler

import (
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant"
	pbmerchantdocument "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant_document"
)

type MerchantQueryHandleGrpc interface {
	pb.MerchantQueryServiceServer
}

type MerchantCommandHandleGrpc interface {
	pb.MerchantCommandServiceServer
}

type MerchantDocumentQueryHandleGrpc interface {
	pbmerchantdocument.MerchantDocumentQueryServiceServer
}

type MerchantDocumentCommandHandleGrpc interface {
	pbmerchantdocument.MerchantDocumentCommandServiceServer
}
