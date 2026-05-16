package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type RestoranHalalRepository interface {
	FindAll(ctx context.Context) ([]domain.RestoranHalal, error)
	FindByID(ctx context.Context, id int) (*domain.RestoranHalal, error)
	FindByKota(ctx context.Context, kota string) ([]domain.RestoranHalal, error)
	Create(ctx context.Context, r *domain.RestoranHalal) error
	Update(ctx context.Context, r *domain.RestoranHalal) error
	Delete(ctx context.Context, id int) error
}

type restoranHalalRepository struct {
	db *pgxpool.Pool
}

func NewRestoranHalalRepository(db *pgxpool.Pool) RestoranHalalRepository {
	return &restoranHalalRepository{db}
}

const restoranHalalColumns = `id, nama_resto, kota, alamat_resto,
	latitude, longitude, ket_resto, harga_resto, foto, created_at`

func scanRestoranHalal(rows pgx.Rows) (domain.RestoranHalal, error) {
	r := domain.RestoranHalal{}
	err := rows.Scan(
		&r.ID, &r.NamaResto, &r.Kota, &r.AlamatResto,
		&r.Latitude, &r.Longitude, &r.KetResto, &r.HargaResto, &r.Foto, &r.CreatedAt,
	)
	return r, err
}

func (r *restoranHalalRepository) FindAll(ctx context.Context) ([]domain.RestoranHalal, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+restoranHalalColumns+` FROM restoran_halal ORDER BY kota`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.RestoranHalal
	for rows.Next() {
		res, err := scanRestoranHalal(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, rows.Err()
}

func (r *restoranHalalRepository) FindByID(ctx context.Context, id int) (*domain.RestoranHalal, error) {
	res := &domain.RestoranHalal{}
	err := r.db.QueryRow(ctx,
		`SELECT `+restoranHalalColumns+` FROM restoran_halal WHERE id = $1`, id,
	).Scan(
		&res.ID, &res.NamaResto, &res.Kota, &res.AlamatResto,
		&res.Latitude, &res.Longitude, &res.KetResto, &res.HargaResto, &res.Foto, &res.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return res, nil
}

func (r *restoranHalalRepository) FindByKota(ctx context.Context, kota string) ([]domain.RestoranHalal, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+restoranHalalColumns+` FROM restoran_halal WHERE kota = $1 ORDER BY nama_resto`, kota,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.RestoranHalal
	for rows.Next() {
		res, err := scanRestoranHalal(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, rows.Err()
}

func (r *restoranHalalRepository) Create(ctx context.Context, res *domain.RestoranHalal) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO restoran_halal
			(nama_resto, kota, alamat_resto, latitude, longitude, ket_resto, harga_resto, foto)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at`,
		res.NamaResto, res.Kota, res.AlamatResto,
		res.Latitude, res.Longitude, res.KetResto, res.HargaResto, res.Foto,
	).Scan(&res.ID, &res.CreatedAt)
}

func (r *restoranHalalRepository) Update(ctx context.Context, res *domain.RestoranHalal) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE restoran_halal
		SET nama_resto = $1, kota = $2, alamat_resto = $3,
		    latitude = $4, longitude = $5, ket_resto = $6, harga_resto = $7, foto = $8
		WHERE id = $9`,
		res.NamaResto, res.Kota, res.AlamatResto,
		res.Latitude, res.Longitude, res.KetResto, res.HargaResto, res.Foto, res.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (r *restoranHalalRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM restoran_halal WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}