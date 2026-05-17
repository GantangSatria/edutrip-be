package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type WisataRepository interface {
	FindAll(ctx context.Context) ([]domain.Wisata, error)
	FindByID(ctx context.Context, id int) (*domain.Wisata, error)
	FindByKota(ctx context.Context, kota string) ([]domain.Wisata, error)
	FindByKotaAndKategori(ctx context.Context, kota, kategori string) ([]domain.Wisata, error)
	Create(ctx context.Context, w *domain.Wisata) error
	Update(ctx context.Context, w *domain.Wisata) error
	Delete(ctx context.Context, id int) error
}

type wisataRepository struct {
	db *pgxpool.Pool
}

func NewWisataRepository(db *pgxpool.Pool) WisataRepository {
	return &wisataRepository{db}
}

const wisataColumns = `id, kota, kategori_wisata, nama_wisata, tiket_wisata,
	alamat_wisata, latitude, longitude, ket_wisata, foto, created_at`

func scanWisata(rows pgx.Rows) (domain.Wisata, error) {
	w := domain.Wisata{}
	err := rows.Scan(
		&w.ID, &w.Kota, &w.KategoriWisata, &w.NamaWisata, &w.TiketWisata,
		&w.AlamatWisata, &w.Latitude, &w.Longitude, &w.KetWisata, &w.Foto, &w.CreatedAt,
	)
	return w, err
}

func (r *wisataRepository) FindAll(ctx context.Context) ([]domain.Wisata, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+wisataColumns+` FROM wisata ORDER BY kota, kategori_wisata`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Wisata
	for rows.Next() {
		w, err := scanWisata(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, w)
	}
	return result, rows.Err()
}

func (r *wisataRepository) FindByID(ctx context.Context, id int) (*domain.Wisata, error) {
	w := &domain.Wisata{}
	err := r.db.QueryRow(ctx,
		`SELECT `+wisataColumns+` FROM wisata WHERE id = $1`, id,
	).Scan(
		&w.ID, &w.Kota, &w.KategoriWisata, &w.NamaWisata, &w.TiketWisata,
		&w.AlamatWisata, &w.Latitude, &w.Longitude, &w.KetWisata, &w.Foto, &w.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return w, nil
}

func (r *wisataRepository) FindByKota(ctx context.Context, kota string) ([]domain.Wisata, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+wisataColumns+` FROM wisata WHERE kota = $1 ORDER BY kategori_wisata`, kota,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Wisata
	for rows.Next() {
		w, err := scanWisata(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, w)
	}
	return result, rows.Err()
}

func (r *wisataRepository) FindByKotaAndKategori(ctx context.Context, kota, kategori string) ([]domain.Wisata, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+wisataColumns+` FROM wisata WHERE kota = $1 AND kategori_wisata = $2 ORDER BY nama_wisata`,
		kota, kategori,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Wisata
	for rows.Next() {
		w, err := scanWisata(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, w)
	}
	return result, rows.Err()
}

func (r *wisataRepository) Create(ctx context.Context, w *domain.Wisata) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO wisata
			(kota, kategori_wisata, nama_wisata, tiket_wisata, alamat_wisata, latitude, longitude, ket_wisata, foto)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at`,
		w.Kota, w.KategoriWisata, w.NamaWisata, w.TiketWisata,
		w.AlamatWisata, w.Latitude, w.Longitude, w.KetWisata, w.Foto,
	).Scan(&w.ID, &w.CreatedAt)
}

func (r *wisataRepository) Update(ctx context.Context, w *domain.Wisata) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE wisata
		SET kota = $1, kategori_wisata = $2, nama_wisata = $3, tiket_wisata = $4,
		    alamat_wisata = $5, latitude = $6, longitude = $7, ket_wisata = $8, foto = $9
		WHERE id = $10`,
		w.Kota, w.KategoriWisata, w.NamaWisata, w.TiketWisata,
		w.AlamatWisata, w.Latitude, w.Longitude, w.KetWisata, w.Foto, w.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (r *wisataRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM wisata WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}