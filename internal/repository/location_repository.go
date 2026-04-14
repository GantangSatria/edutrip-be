package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/internal/domain"
)

type LocationRepository interface {
	FindAll(ctx context.Context) ([]domain.Location, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Location, error)
	Create(ctx context.Context, loc *domain.Location) error
	Update(ctx context.Context, loc *domain.Location) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type locationRepository struct {
	db *pgxpool.Pool
}

func NewLocationRepository(db *pgxpool.Pool) LocationRepository {
	return &locationRepository{db}
}

func (r *locationRepository) FindAll(ctx context.Context) ([]domain.Location, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, admin_id, name, city, category, address, price,
		        latitude, longitude, description, website, is_halal, created_at
		 FROM locations ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locs []domain.Location
	for rows.Next() {
		var l domain.Location
		if err := rows.Scan(&l.ID, &l.AdminID, &l.Name, &l.City, &l.Category,
			&l.Address, &l.Price, &l.Latitude, &l.Longitude,
			&l.Description, &l.Website, &l.IsHalal, &l.CreatedAt); err != nil {
			return nil, err
		}
		locs = append(locs, l)
	}
	return locs, nil
}

func (r *locationRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Location, error) {
	l := &domain.Location{}
	err := r.db.QueryRow(ctx,
		`SELECT id, admin_id, name, city, category, address, price,
		        latitude, longitude, description, website, is_halal, created_at
		 FROM locations WHERE id = $1`, id,
	).Scan(&l.ID, &l.AdminID, &l.Name, &l.City, &l.Category,
		&l.Address, &l.Price, &l.Latitude, &l.Longitude,
		&l.Description, &l.Website, &l.IsHalal, &l.CreatedAt)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (r *locationRepository) Create(ctx context.Context, l *domain.Location) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO locations (admin_id, name, city, category, address, price,
		                        latitude, longitude, description, website, is_halal)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		 RETURNING id, created_at`,
		l.AdminID, l.Name, l.City, l.Category, l.Address, l.Price,
		l.Latitude, l.Longitude, l.Description, l.Website, l.IsHalal,
	).Scan(&l.ID, &l.CreatedAt)
}

func (r *locationRepository) Update(ctx context.Context, l *domain.Location) error {
	_, err := r.db.Exec(ctx,
		`UPDATE locations SET name=$1, city=$2, category=$3, address=$4, price=$5,
		                      latitude=$6, longitude=$7, description=$8, website=$9, is_halal=$10
		 WHERE id=$11`,
		l.Name, l.City, l.Category, l.Address, l.Price,
		l.Latitude, l.Longitude, l.Description, l.Website, l.IsHalal, l.ID,
	)
	return err
}

func (r *locationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM locations WHERE id = $1`, id)
	return err
}