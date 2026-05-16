package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type FasilitasIbadahRepository interface {
	FindAll(ctx context.Context) ([]domain.FasilitasIbadah, error)
	FindByID(ctx context.Context, id int) (*domain.FasilitasIbadah, error)
	FindByKota(ctx context.Context, kota string) ([]domain.FasilitasIbadah, error)
	Create(ctx context.Context, f *domain.FasilitasIbadah) error
	Update(ctx context.Context, f *domain.FasilitasIbadah) error
	Delete(ctx context.Context, id int) error
}

type fasilitasIbadahRepository struct {
	db *pgxpool.Pool
}

func NewFasilitasIbadahRepository(db *pgxpool.Pool) FasilitasIbadahRepository {
	return &fasilitasIbadahRepository{db}
}

const fasilitasIbadahColumns = `id, kota, tipe_fas, nama_fas_ibadah, lokasi_fas_ibadah,
	latitude, longitude, foto, created_at`

func scanFasilitasIbadah(rows pgx.Rows) (domain.FasilitasIbadah, error) {
	f := domain.FasilitasIbadah{}
	err := rows.Scan(
		&f.ID, &f.Kota, &f.TipeFas, &f.NamaFasIbadah, &f.LokasiFasIbadah,
		&f.Latitude, &f.Longitude, &f.Foto, &f.CreatedAt,
	)
	return f, err
}

func (r *fasilitasIbadahRepository) FindAll(ctx context.Context) ([]domain.FasilitasIbadah, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+fasilitasIbadahColumns+` FROM fasilitas_ibadah ORDER BY kota, tipe_fas`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.FasilitasIbadah
	for rows.Next() {
		f, err := scanFasilitasIbadah(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, rows.Err()
}

func (r *fasilitasIbadahRepository) FindByID(ctx context.Context, id int) (*domain.FasilitasIbadah, error) {
	f := &domain.FasilitasIbadah{}
	err := r.db.QueryRow(ctx,
		`SELECT `+fasilitasIbadahColumns+` FROM fasilitas_ibadah WHERE id = $1`, id,
	).Scan(
		&f.ID, &f.Kota, &f.TipeFas, &f.NamaFasIbadah, &f.LokasiFasIbadah,
		&f.Latitude, &f.Longitude, &f.Foto, &f.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return f, nil
}

func (r *fasilitasIbadahRepository) FindByKota(ctx context.Context, kota string) ([]domain.FasilitasIbadah, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+fasilitasIbadahColumns+` FROM fasilitas_ibadah WHERE kota = $1 ORDER BY tipe_fas`, kota,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.FasilitasIbadah
	for rows.Next() {
		f, err := scanFasilitasIbadah(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, rows.Err()
}

func (r *fasilitasIbadahRepository) Create(ctx context.Context, f *domain.FasilitasIbadah) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO fasilitas_ibadah
			(kota, tipe_fas, nama_fas_ibadah, lokasi_fas_ibadah, latitude, longitude, foto)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at`,
		f.Kota, f.TipeFas, f.NamaFasIbadah, f.LokasiFasIbadah,
		f.Latitude, f.Longitude, f.Foto,
	).Scan(&f.ID, &f.CreatedAt)
}

func (r *fasilitasIbadahRepository) Update(ctx context.Context, f *domain.FasilitasIbadah) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE fasilitas_ibadah
		SET kota = $1, tipe_fas = $2, nama_fas_ibadah = $3, lokasi_fas_ibadah = $4,
		    latitude = $5, longitude = $6, foto = $7
		WHERE id = $8`,
		f.Kota, f.TipeFas, f.NamaFasIbadah, f.LokasiFasIbadah,
		f.Latitude, f.Longitude, f.Foto, f.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (r *fasilitasIbadahRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM fasilitas_ibadah WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}