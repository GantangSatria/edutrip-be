package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/internal/domain"
)

type TripRepository interface {
	FindAll(ctx context.Context) ([]domain.Trip, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Trip, error)
	Create(ctx context.Context, trip *domain.Trip) error
	Update(ctx context.Context, trip *domain.Trip) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type tripRepository struct {
	db *pgxpool.Pool
}

func NewTripRepository(db *pgxpool.Pool) TripRepository {
	return &tripRepository{db}
}

func (r *tripRepository) FindAll(ctx context.Context) ([]domain.Trip, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, admin_id, title, description, city, theme, duration_days,
		        price, facilities, itinerary_summary, status, created_at
		 FROM trips ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []domain.Trip
	for rows.Next() {
		var t domain.Trip
		if err := rows.Scan(&t.ID, &t.AdminID, &t.Title, &t.Description, &t.City,
			&t.Theme, &t.DurationDays, &t.Price, &t.Facilities,
			&t.ItinerarySummary, &t.Status, &t.CreatedAt); err != nil {
			return nil, err
		}
		trips = append(trips, t)
	}
	return trips, nil
}

func (r *tripRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	t := &domain.Trip{}
	err := r.db.QueryRow(ctx,
		`SELECT id, admin_id, title, description, city, theme, duration_days,
		        price, facilities, itinerary_summary, status, created_at
		 FROM trips WHERE id = $1`, id,
	).Scan(&t.ID, &t.AdminID, &t.Title, &t.Description, &t.City,
		&t.Theme, &t.DurationDays, &t.Price, &t.Facilities,
		&t.ItinerarySummary, &t.Status, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *tripRepository) Create(ctx context.Context, t *domain.Trip) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO trips (admin_id, title, description, city, theme, duration_days,
		                    price, facilities, itinerary_summary, status)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		 RETURNING id, created_at`,
		t.AdminID, t.Title, t.Description, t.City, t.Theme, t.DurationDays,
		t.Price, t.Facilities, t.ItinerarySummary, t.Status,
	).Scan(&t.ID, &t.CreatedAt)
}

func (r *tripRepository) Update(ctx context.Context, t *domain.Trip) error {
	_, err := r.db.Exec(ctx,
		`UPDATE trips SET title=$1, description=$2, city=$3, theme=$4, duration_days=$5,
		                  price=$6, facilities=$7, itinerary_summary=$8, status=$9
		 WHERE id=$10`,
		t.Title, t.Description, t.City, t.Theme, t.DurationDays,
		t.Price, t.Facilities, t.ItinerarySummary, t.Status, t.ID,
	)
	return err
}

func (r *tripRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM trips WHERE id = $1`, id)
	return err
}