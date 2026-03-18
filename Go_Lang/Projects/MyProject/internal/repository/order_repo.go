package repository

import (
	"context"
	"database/sql"
	"errors"

	"MyProject/internal/models"
	"MyProject/internal/database"
)

var (
	ErrOrderExists = errors.New("Order ID already exists")
)

type OrderRepo interface {
	GetAll(ctx context.Context) ([]models.Order, error)
	GetByID(ctx context.Context, id string) (*models.Order, error)
	Create(ctx context.Context, order models.Order) error
}

type orderRepo struct {
	database *sql.DB
}

func NewOrderRepo(db *sql.DB) OrderRepo {
	return &orderRepo{
		database: db,
	}
}

func (r *orderRepo) GetAll(ctx context.Context) ([]models.Order, error) {
	query := `SELECT id, facility_code, status, created_at FROM orders`
	rows, err := r.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var orders []models.Order

	for rows.Next() {
		var o models.Order
		rows.Scan(&o.ID, &o.FacilityCode, &o.Status, &o.CreatedAt)
		orders = append(orders, o)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepo) GetByID(ctx context.Context, id string) (*models.Order, error) {
	query := `SELECT id, facility_code, status, created_at FROM orders WHERE id=?`

	var o models.Order
	err := r.database.QueryRowContext(ctx, query, id).Scan(
		&o.ID,
		&o.FacilityCode,
		&o.Status,
		&o.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &o, nil
}

func (r *orderRepo) Create(ctx context.Context, order models.Order) error {
	query := `INSERT INTO orders (id, facility_code, status, created_at) VALUES (?,?,?,?)`
	_, err := r.database.ExecContext(ctx, query,
		order.ID,
		order.FacilityCode,
		order.Status,
		order.CreatedAt,
	)

	if database.IsDuplicateError(err) {
		return ErrOrderExists
	}

	return err
}