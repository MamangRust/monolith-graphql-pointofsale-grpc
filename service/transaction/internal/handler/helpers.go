package handler

import (
	pbutils "github.com/MamangRust/monolith-graphql-pointofsale-pb/api"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb/transaction"
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

func mapResponseTransaction(transaction *db.Transaction) *pb.TransactionResponse {
	if transaction == nil {
		return nil
	}
	var createdAtStr string
	if transaction.CreatedAt.Valid {
		createdAtStr = transaction.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var updatedAtStr string
	if transaction.UpdatedAt.Valid {
		updatedAtStr = transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var changeAmount int32
	if transaction.ChangeAmount != nil {
		changeAmount = *transaction.ChangeAmount
	}
	return &pb.TransactionResponse{
		Id:            transaction.TransactionID,
		OrderId:       transaction.OrderID,
		MerchantId:    transaction.MerchantID,
		PaymentMethod: transaction.PaymentMethod,
		Amount:        transaction.Amount,
		ChangeAmount:  changeAmount,
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     createdAtStr,
		UpdatedAt:     updatedAtStr,
	}
}

func mapResponsesTransaction(transactions []*db.GetTransactionsRow) []*pb.TransactionResponse {
	var mappedTransactions []*pb.TransactionResponse
	for _, t := range transactions {
		if t == nil {
			continue
		}
		var createdAtStr string
		if t.CreatedAt.Valid {
			createdAtStr = t.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var updatedAtStr string
		if t.UpdatedAt.Valid {
			updatedAtStr = t.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var changeAmount int32
		if t.ChangeAmount != nil {
			changeAmount = *t.ChangeAmount
		}
		mappedTransactions = append(mappedTransactions, &pb.TransactionResponse{
			Id:            t.TransactionID,
			OrderId:       t.OrderID,
			MerchantId:    t.MerchantID,
			PaymentMethod: t.PaymentMethod,
			Amount:        t.Amount,
			ChangeAmount:  changeAmount,
			PaymentStatus: t.PaymentStatus,
			CreatedAt:     createdAtStr,
			UpdatedAt:     updatedAtStr,
		})
	}
	return mappedTransactions
}

func mapResponsesTransactionByMerchant(transactions []*db.GetTransactionByMerchantRow) []*pb.TransactionResponse {
	var mappedTransactions []*pb.TransactionResponse
	for _, t := range transactions {
		if t == nil {
			continue
		}
		var createdAtStr string
		if t.CreatedAt.Valid {
			createdAtStr = t.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var updatedAtStr string
		if t.UpdatedAt.Valid {
			updatedAtStr = t.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var changeAmount int32
		if t.ChangeAmount != nil {
			changeAmount = *t.ChangeAmount
		}
		mappedTransactions = append(mappedTransactions, &pb.TransactionResponse{
			Id:            t.TransactionID,
			OrderId:       t.OrderID,
			MerchantId:    t.MerchantID,
			PaymentMethod: t.PaymentMethod,
			Amount:        t.Amount,
			ChangeAmount:  changeAmount,
			PaymentStatus: t.PaymentStatus,
			CreatedAt:     createdAtStr,
			UpdatedAt:     updatedAtStr,
		})
	}
	return mappedTransactions
}

func mapResponseTransactionDeleteAt(transaction *db.Transaction) *pb.TransactionResponseDeleteAt {
	if transaction == nil {
		return nil
	}
	var createdAtStr string
	if transaction.CreatedAt.Valid {
		createdAtStr = transaction.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var updatedAtStr string
	if transaction.UpdatedAt.Valid {
		updatedAtStr = transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	var deletedAt *wrapperspb.StringValue
	if transaction.DeletedAt.Valid {
		deletedAt = wrapperspb.String(transaction.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	}
	var changeAmount int32
	if transaction.ChangeAmount != nil {
		changeAmount = *transaction.ChangeAmount
	}

	return &pb.TransactionResponseDeleteAt{
		Id:            transaction.TransactionID,
		OrderId:       transaction.OrderID,
		MerchantId:    transaction.MerchantID,
		PaymentMethod: transaction.PaymentMethod,
		Amount:        transaction.Amount,
		ChangeAmount:  changeAmount,
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     createdAtStr,
		UpdatedAt:     updatedAtStr,
		DeletedAt:     deletedAt,
	}
}

func mapResponsesTransactionActive(transactions []*db.GetTransactionsActiveRow) []*pb.TransactionResponseDeleteAt {
	var mappedTransactions []*pb.TransactionResponseDeleteAt
	for _, t := range transactions {
		if t == nil {
			continue
		}
		var createdAtStr string
		if t.CreatedAt.Valid {
			createdAtStr = t.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var updatedAtStr string
		if t.UpdatedAt.Valid {
			updatedAtStr = t.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAt *wrapperspb.StringValue
		if t.DeletedAt.Valid {
			deletedAt = wrapperspb.String(t.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}
		var changeAmount int32
		if t.ChangeAmount != nil {
			changeAmount = *t.ChangeAmount
		}
		mappedTransactions = append(mappedTransactions, &pb.TransactionResponseDeleteAt{
			Id:            t.TransactionID,
			OrderId:       t.OrderID,
			MerchantId:    t.MerchantID,
			PaymentMethod: t.PaymentMethod,
			Amount:        t.Amount,
			ChangeAmount:  changeAmount,
			PaymentStatus: t.PaymentStatus,
			CreatedAt:     createdAtStr,
			UpdatedAt:     updatedAtStr,
			DeletedAt:     deletedAt,
		})
	}
	return mappedTransactions
}

func mapResponsesTransactionTrashed(transactions []*db.GetTransactionsTrashedRow) []*pb.TransactionResponseDeleteAt {
	var mappedTransactions []*pb.TransactionResponseDeleteAt
	for _, t := range transactions {
		if t == nil {
			continue
		}
		var createdAtStr string
		if t.CreatedAt.Valid {
			createdAtStr = t.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var updatedAtStr string
		if t.UpdatedAt.Valid {
			updatedAtStr = t.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		}
		var deletedAt *wrapperspb.StringValue
		if t.DeletedAt.Valid {
			deletedAt = wrapperspb.String(t.DeletedAt.Time.Format("2006-01-02 15:04:05"))
		}
		var changeAmount int32
		if t.ChangeAmount != nil {
			changeAmount = *t.ChangeAmount
		}
		mappedTransactions = append(mappedTransactions, &pb.TransactionResponseDeleteAt{
			Id:            t.TransactionID,
			OrderId:       t.OrderID,
			MerchantId:    t.MerchantID,
			PaymentMethod: t.PaymentMethod,
			Amount:        t.Amount,
			ChangeAmount:  changeAmount,
			PaymentStatus: t.PaymentStatus,
			CreatedAt:     createdAtStr,
			UpdatedAt:     updatedAtStr,
			DeletedAt:     deletedAt,
		})
	}
	return mappedTransactions
}

func mapResponseTransactionMonthAmountSuccess(row *db.GetMonthlyAmountTransactionSuccessRow) *pb.TransactionMonthlyAmountSuccess {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyAmountSuccess{
		Year:         row.Year,
		Month:        row.Month,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyAmountSuccess(rows []*db.GetMonthlyAmountTransactionSuccessRow) []*pb.TransactionMonthlyAmountSuccess {
	var transaction []*pb.TransactionMonthlyAmountSuccess
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthAmountSuccess(row))
	}
	return transaction
}

func mapResponseTransactionYearAmountSuccess(row *db.GetYearlyAmountTransactionSuccessRow) *pb.TransactionYearlyAmountSuccess {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyAmountSuccess{
		Year:         row.Year,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyAmountSuccess(rows []*db.GetYearlyAmountTransactionSuccessRow) []*pb.TransactionYearlyAmountSuccess {
	var transaction []*pb.TransactionYearlyAmountSuccess
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearAmountSuccess(row))
	}
	return transaction
}

func mapResponseTransactionMonthAmountFailed(row *db.GetMonthlyAmountTransactionFailedRow) *pb.TransactionMonthlyAmountFailed {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyAmountFailed{
		Year:        row.Year,
		Month:       row.Month,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyAmountFailed(rows []*db.GetMonthlyAmountTransactionFailedRow) []*pb.TransactionMonthlyAmountFailed {
	var transaction []*pb.TransactionMonthlyAmountFailed
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthAmountFailed(row))
	}
	return transaction
}

func mapResponseTransactionYearAmountFailed(row *db.GetYearlyAmountTransactionFailedRow) *pb.TransactionYearlyAmountFailed {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyAmountFailed{
		Year:        row.Year,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyAmountFailed(rows []*db.GetYearlyAmountTransactionFailedRow) []*pb.TransactionYearlyAmountFailed {
	var transaction []*pb.TransactionYearlyAmountFailed
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearAmountFailed(row))
	}
	return transaction
}

func mapResponseTransactionMonthAmountSuccessByMerchant(row *db.GetMonthlyAmountTransactionSuccessByMerchantRow) *pb.TransactionMonthlyAmountSuccess {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyAmountSuccess{
		Year:         row.Year,
		Month:        row.Month,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyAmountSuccessByMerchant(rows []*db.GetMonthlyAmountTransactionSuccessByMerchantRow) []*pb.TransactionMonthlyAmountSuccess {
	var transaction []*pb.TransactionMonthlyAmountSuccess
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthAmountSuccessByMerchant(row))
	}
	return transaction
}

func mapResponseTransactionYearAmountSuccessByMerchant(row *db.GetYearlyAmountTransactionSuccessByMerchantRow) *pb.TransactionYearlyAmountSuccess {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyAmountSuccess{
		Year:         row.Year,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyAmountSuccessByMerchant(rows []*db.GetYearlyAmountTransactionSuccessByMerchantRow) []*pb.TransactionYearlyAmountSuccess {
	var transaction []*pb.TransactionYearlyAmountSuccess
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearAmountSuccessByMerchant(row))
	}
	return transaction
}

func mapResponseTransactionMonthAmountFailedByMerchant(row *db.GetMonthlyAmountTransactionFailedByMerchantRow) *pb.TransactionMonthlyAmountFailed {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyAmountFailed{
		Year:        row.Year,
		Month:       row.Month,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyAmountFailedByMerchant(rows []*db.GetMonthlyAmountTransactionFailedByMerchantRow) []*pb.TransactionMonthlyAmountFailed {
	var transaction []*pb.TransactionMonthlyAmountFailed
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthAmountFailedByMerchant(row))
	}
	return transaction
}

func mapResponseTransactionYearAmountFailedByMerchant(row *db.GetYearlyAmountTransactionFailedByMerchantRow) *pb.TransactionYearlyAmountFailed {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyAmountFailed{
		Year:        row.Year,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyAmountFailedByMerchant(rows []*db.GetYearlyAmountTransactionFailedByMerchantRow) []*pb.TransactionYearlyAmountFailed {
	var transaction []*pb.TransactionYearlyAmountFailed
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearAmountFailedByMerchant(row))
	}
	return transaction
}

func mapResponseTransactionMonthMethodSuccess(row *db.GetMonthlyTransactionMethodsSuccessRow) *pb.TransactionMonthlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyMethod{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyMethodSuccess(rows []*db.GetMonthlyTransactionMethodsSuccessRow) []*pb.TransactionMonthlyMethod {
	var transaction []*pb.TransactionMonthlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthMethodSuccess(row))
	}
	return transaction
}

func mapResponseTransactionYearlyMethodSuccess(row *db.GetYearlyTransactionMethodsSuccessRow) *pb.TransactionYearlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyMethod{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyMethodSuccess(rows []*db.GetYearlyTransactionMethodsSuccessRow) []*pb.TransactionYearlyMethod {
	var transaction []*pb.TransactionYearlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearlyMethodSuccess(row))
	}
	return transaction
}

func mapResponseTransactionMonthMethodFailed(row *db.GetMonthlyTransactionMethodsFailedRow) *pb.TransactionMonthlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyMethod{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyMethodFailed(rows []*db.GetMonthlyTransactionMethodsFailedRow) []*pb.TransactionMonthlyMethod {
	var transaction []*pb.TransactionMonthlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthMethodFailed(row))
	}
	return transaction
}

func mapResponseTransactionYearlyMethodFailed(row *db.GetYearlyTransactionMethodsFailedRow) *pb.TransactionYearlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyMethod{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyMethodFailed(rows []*db.GetYearlyTransactionMethodsFailedRow) []*pb.TransactionYearlyMethod {
	var transaction []*pb.TransactionYearlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearlyMethodFailed(row))
	}
	return transaction
}

func mapResponseTransactionMonthMethodByMerchantSuccess(row *db.GetMonthlyTransactionMethodsByMerchantSuccessRow) *pb.TransactionMonthlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyMethod{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyMethodByMerchantSuccess(rows []*db.GetMonthlyTransactionMethodsByMerchantSuccessRow) []*pb.TransactionMonthlyMethod {
	var transaction []*pb.TransactionMonthlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthMethodByMerchantSuccess(row))
	}
	return transaction
}

func mapResponseTransactionYearlyMethodByMerchantSuccess(row *db.GetYearlyTransactionMethodsByMerchantSuccessRow) *pb.TransactionYearlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyMethod{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyMethodByMerchantSuccess(rows []*db.GetYearlyTransactionMethodsByMerchantSuccessRow) []*pb.TransactionYearlyMethod {
	var transaction []*pb.TransactionYearlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearlyMethodByMerchantSuccess(row))
	}
	return transaction
}

func mapResponseTransactionMonthMethodByMerchantFailed(row *db.GetMonthlyTransactionMethodsByMerchantFailedRow) *pb.TransactionMonthlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionMonthlyMethod{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionMonthlyMethodByMerchantFailed(rows []*db.GetMonthlyTransactionMethodsByMerchantFailedRow) []*pb.TransactionMonthlyMethod {
	var transaction []*pb.TransactionMonthlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionMonthMethodByMerchantFailed(row))
	}
	return transaction
}

func mapResponseTransactionYearlyMethodByMerchantFailed(row *db.GetYearlyTransactionMethodsByMerchantFailedRow) *pb.TransactionYearlyMethod {
	if row == nil {
		return nil
	}
	return &pb.TransactionYearlyMethod{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func mapResponsesTransactionYearlyMethodByMerchantFailed(rows []*db.GetYearlyTransactionMethodsByMerchantFailedRow) []*pb.TransactionYearlyMethod {
	var transaction []*pb.TransactionYearlyMethod
	for _, row := range rows {
		transaction = append(transaction, mapResponseTransactionYearlyMethodByMerchantFailed(row))
	}
	return transaction
}
