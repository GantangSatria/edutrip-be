package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokoOlehOlehRepository interface {
	FindAll(ctx context.Context) ([]domain.TokoOlehOleh, error)
	FindByID(ctx context.Context, id int) (*domain.TokoOlehOleh, error)
	FindByKota(ctx context.Context, kota string) ([]domain.TokoOlehOleh, error)
}

type tokoOlehOlehRepository struct {
	db *pgxpool.Pool
}

func NewTokoOlehOlehRepository(db *pgxpool.Pool) TokoOlehOlehRepository {
	return &tokoOlehOlehRepository{db}
}

const tokoOlehOlehColumns = `id, kota, nama_belanja, jenis_belanja, ket_belanja,
	alamat_belanja, latitude, longitude, foto, created_at`

func scanTokoOlehOleh(rows pgx.Rows) (domain.TokoOlehOleh, error) {
	t := domain.TokoOlehOleh{}
	err := rows.Scan(
		&t.ID, &t.Kota, &t.NamaBelanja, &t.JenisBelanja, &t.KetBelanja,
		&t.AlamatBelanja, &t.Latitude, &t.Longitude, &t.Foto, &t.CreatedAt,
	)
	return t, err
}

func (r *tokoOlehOlehRepository) FindAll(ctx context.Context) ([]domain.TokoOlehOleh, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+tokoOlehOlehColumns+` FROM toko_oleh_oleh ORDER BY kota`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.TokoOlehOleh
	for rows.Next() {
		t, err := scanTokoOlehOleh(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, rows.Err()
}

func (r *tokoOlehOlehRepository) FindByID(ctx context.Context, id int) (*domain.TokoOlehOleh, error) {
	t := &domain.TokoOlehOleh{}
	err := r.db.QueryRow(ctx,
		`SELECT `+tokoOlehOlehColumns+` FROM toko_oleh_oleh WHERE id = $1`, id,
	).Scan(
		&t.ID, &t.Kota, &t.NamaBelanja, &t.JenisBelanja, &t.KetBelanja,
		&t.AlamatBelanja, &t.Latitude, &t.Longitude, &t.Foto, &t.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return t, nil
}

func (r *tokoOlehOlehRepository) FindByKota(ctx context.Context, kota string) ([]domain.TokoOlehOleh, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+tokoOlehOlehColumns+` FROM toko_oleh_oleh WHERE kota = $1 ORDER BY nama_belanja`, kota,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.TokoOlehOleh
	for rows.Next() {
		t, err := scanTokoOlehOleh(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, rows.Err()
}