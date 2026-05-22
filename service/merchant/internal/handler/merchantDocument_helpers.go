package handler

import (
	db "github.com/MamangRust/monolith-point-of-sale-pkg/database/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/merchant_document"
)

// Map helpers
func mapMerchantDocument(doc *db.MerchantDocument) *pb.MerchantDocument {
	if doc == nil {
		return nil
	}
	return &pb.MerchantDocument{
		DocumentId:   int32(doc.DocumentID),
		MerchantId:   int32(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         mapSqlNullString(doc.Note),
		UploadedAt:   mapSqlNullTime(doc.CreatedAt),
		UpdatedAt:    mapSqlNullTime(doc.UpdatedAt),
	}
}

func mapResponsesGetMerchantDocumentsRow(docs []*db.GetMerchantDocumentsRow) []*pb.MerchantDocument {
	var res []*pb.MerchantDocument
	for _, doc := range docs {
		res = append(res, &pb.MerchantDocument{
			DocumentId:   int32(doc.DocumentID),
			MerchantId:   int32(doc.MerchantID),
			DocumentType: doc.DocumentType,
			DocumentUrl:  doc.DocumentUrl,
			Status:       doc.Status,
			Note:         mapSqlNullString(doc.Note),
			UploadedAt:   mapSqlNullTime(doc.CreatedAt),
			UpdatedAt:    mapSqlNullTime(doc.UpdatedAt),
		})
	}
	return res
}

func mapResponsesGetActiveMerchantDocumentsRow(docs []*db.GetActiveMerchantDocumentsRow) []*pb.MerchantDocument {
	var res []*pb.MerchantDocument
	for _, doc := range docs {
		res = append(res, &pb.MerchantDocument{
			DocumentId:   int32(doc.DocumentID),
			MerchantId:   int32(doc.MerchantID),
			DocumentType: doc.DocumentType,
			DocumentUrl:  doc.DocumentUrl,
			Status:       doc.Status,
			Note:         mapSqlNullString(doc.Note),
			UploadedAt:   mapSqlNullTime(doc.CreatedAt),
			UpdatedAt:    mapSqlNullTime(doc.UpdatedAt),
		})
	}
	return res
}

func mapMerchantDocumentDeleteAt(doc *db.MerchantDocument) *pb.MerchantDocumentDeleteAt {
	if doc == nil {
		return nil
	}
	var deletedAt *wrapperspb.StringValue
	if doc.DeletedAt.Valid {
		deletedAt = wrapperspb.String(doc.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}

	return &pb.MerchantDocumentDeleteAt{
		DocumentId:   int32(doc.DocumentID),
		MerchantId:   int32(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         mapSqlNullString(doc.Note),
		UploadedAt:   mapSqlNullTime(doc.CreatedAt),
		UpdatedAt:    mapSqlNullTime(doc.UpdatedAt),
		DeletedAt:    deletedAt,
	}
}

func mapResponsesGetTrashedMerchantDocumentsRow(docs []*db.GetTrashedMerchantDocumentsRow) []*pb.MerchantDocumentDeleteAt {
	var res []*pb.MerchantDocumentDeleteAt
	for _, doc := range docs {
		var deletedAt *wrapperspb.StringValue
		if doc.DeletedAt.Valid {
			deletedAt = wrapperspb.String(doc.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}
		res = append(res, &pb.MerchantDocumentDeleteAt{
			DocumentId:   int32(doc.DocumentID),
			MerchantId:   int32(doc.MerchantID),
			DocumentType: doc.DocumentType,
			DocumentUrl:  doc.DocumentUrl,
			Status:       doc.Status,
			Note:         mapSqlNullString(doc.Note),
			UploadedAt:   mapSqlNullTime(doc.CreatedAt),
			UpdatedAt:    mapSqlNullTime(doc.UpdatedAt),
			DeletedAt:    deletedAt,
		})
	}
	return res
}
