package repository

import (
	"context"
	"errors"
	"database/sql"
	
	"MyProject/internal/models"
	"MyProject/internal/database"
)

var (
	ErrFacilityExists = errors.New("Facility ID already exists")
)

type FacilityRepo interface {
	GetAll(ctx context.Context) ([]models.Facility, error)
	GetByCode(ctx context.Context, code string) (*models.Facility, error)
	Create(ctx context.Context, facility models.Facility) error
}

type facilityRepo struct {
	db *sql.DB
}

func NewFacilityRepo(database *sql.DB) FacilityRepo {
	return &facilityRepo{
		db: database,
	}
}

func (r *facilityRepo) GetAll(ctx context.Context) ([]models.Facility, error) {
	query := `SELECT * FROM facilities`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	facilities := []models.Facility{}

	for rows.Next() {
		var f models.Facility
		if err := rows.Scan(&f.Code, &f.Name, &f.Address); err != nil {
			return nil, err
		}
		facilities = append(facilities, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return facilities, nil
}
func (r *facilityRepo) GetByCode(ctx context.Context, code string) (*models.Facility, error) {
	query := `SELECT * FROM facilities WHERE code=?`
	var f models.Facility
	err := r.db.QueryRowContext(ctx, query, code).Scan(&f.Code, &f.Name, &f.Address)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func(r *facilityRepo) Create(ctx context.Context, f models.Facility) error {
	query := `INSERT INTO facilities (code, name, address) VALUES  (?,?,?)`
	_, err := r.db.ExecContext(ctx, query,
	f.Code,
	f.Name,
	f.Address,
	)
	
	if database.IsDuplicateError(err) {
		return ErrFacilityExists
	}
	return err
}