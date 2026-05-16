package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type TokoOlehOlehRepository interface {
	FindAll(ctx context.Context) ([]domain.TokoOlehOleh, error)
	FindByID(ctx context.Context, id int) (*domain.TokoOlehOleh, error)
	FindByKota(ctx context.Context, kota string) ([]domain.TokoOlehOleh, error)
	Create(ctx context.Context, t *domain.TokoOlehOleh) error
	Update(ctx context.Context, t *domain.TokoOlehOleh) error
	Delete(ctx context.Context, id int) error
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
			return nil, utils.ErrNotFound
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

func (r *tokoOlehOlehRepository) Create(ctx context.Context, t *domain.TokoOlehOleh) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO toko_oleh_oleh
			(kota, nama_belanja, jenis_belanja, ket_belanja, alamat_belanja, latitude, longitude, foto)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at`,
		t.Kota, t.NamaBelanja, t.JenisBelanja, t.KetBelanja,
		t.AlamatBelanja, t.Latitude, t.Longitude, t.Foto,
	).Scan(&t.ID, &t.CreatedAt)
}

func (r *tokoOlehOlehRepository) Update(ctx context.Context, t *domain.TokoOlehOleh) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE toko_oleh_oleh
		SET kota = $1, nama_belanja = $2, jenis_belanja = $3, ket_belanja = $4,
		    alamat_belanja = $5, latitude = $6, longitude = $7, foto = $8
		WHERE id = $9`,
		t.Kota, t.NamaBelanja, t.JenisBelanja, t.KetBelanja,
		t.AlamatBelanja, t.Latitude, t.Longitude, t.Foto, t.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (r *tokoOlehOlehRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM toko_oleh_oleh WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}