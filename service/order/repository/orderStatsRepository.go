package repository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/errors/order_errors"
	"github.com/jackc/pgx/v5/pgtype"
)

type orderStatsRepository struct {
	db *db.Queries
}

func NewOrderStatsRepository(db *db.Queries) OrderStatsRepository {
	return &orderStatsRepository{
		db: db,
	}
}

func (r *orderStatsRepository) GetMonthlyTotalRevenue(ctx context.Context, req *requests.MonthTotalRevenue) ([]*db.GetMonthlyTotalRevenueRow, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalRevenue(ctx, db.GetMonthlyTotalRevenueParams{
		Extract:     pgtype.Date{Time: currentMonthStart, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: pgtype.Timestamp{Time: prevMonthStart, Valid: true},
		CreatedAt_3: pgtype.Timestamp{Time: prevMonthEnd, Valid: true},
	})
	if err != nil {
		return nil, order_errors.ErrGetMonthlyTotalRevenue
	}

	return res, nil
}

func (r *orderStatsRepository) GetYearlyTotalRevenue(ctx context.Context, year int) ([]*db.GetYearlyTotalRevenueRow, error) {
	res, err := r.db.GetYearlyTotalRevenue(ctx, int32(year))
	if err != nil {
		return nil, order_errors.ErrGetYearlyTotalRevenue
	}

	return res, nil
}

func (r *orderStatsRepository) GetMonthlyTotalRevenueById(ctx context.Context, req *requests.MonthTotalRevenueOrder) ([]*db.GetMonthlyTotalRevenueByIdRow, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalRevenueById(ctx, db.GetMonthlyTotalRevenueByIdParams{
		Extract:     pgtype.Date{Time: currentMonthStart, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: pgtype.Timestamp{Time: prevMonthStart, Valid: true},
		CreatedAt_3: pgtype.Timestamp{Time: prevMonthEnd, Valid: true},
		OrderID:     int32(req.OrderID),
	})
	if err != nil {
		return nil, order_errors.ErrGetMonthlyTotalRevenueById
	}

	return res, nil
}

func (r *orderStatsRepository) GetYearlyTotalRevenueById(ctx context.Context, req *requests.YearTotalRevenueOrder) ([]*db.GetYearlyTotalRevenueByIdRow, error) {
	res, err := r.db.GetYearlyTotalRevenueById(ctx, db.GetYearlyTotalRevenueByIdParams{
		Column1: int32(req.Year),
		OrderID: int32(req.OrderID),
	})
	if err != nil {
		return nil, order_errors.ErrGetYearlyTotalRevenueById
	}

	return res, nil
}

func (r *orderStatsRepository) GetMonthlyOrder(ctx context.Context, year int) ([]*db.GetMonthlyOrderRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	res, err := r.db.GetMonthlyOrder(ctx, yearStart)
	if err != nil {
		return nil, order_errors.ErrGetMonthlyOrder
	}

	return res, nil
}

func (r *orderStatsRepository) GetYearlyOrder(ctx context.Context, year int) ([]*db.GetYearlyOrderRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	res, err := r.db.GetYearlyOrder(ctx, yearStart)
	if err != nil {
		return nil, order_errors.ErrGetYearlyOrder
	}

	return res, nil
}
