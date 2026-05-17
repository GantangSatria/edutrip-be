package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type HotelRepository interface {
	FindAll(ctx context.Context) ([]domain.Hotel, error)
	FindByID(ctx context.Context, id int) (*domain.Hotel, error)
	FindByKota(ctx context.Context, kota string) ([]domain.Hotel, error)
	FindByKotaAndTipe(ctx context.Context, kota, tipe string) ([]domain.Hotel, error)
	Create(ctx context.Context, h *domain.Hotel) error
	Update(ctx context.Context, h *domain.Hotel) error
	Delete(ctx context.Context, id int) error
}

type hotelRepository struct {
	db *pgxpool.Pool
}

func NewHotelRepository(db *pgxpool.Pool) HotelRepository {
	return &hotelRepository{db}
}

const hotelColumns = `id, kota, nama_hotel, tipe_hotel, harga_hotel,
	alamat_hotel, latitude, longitude, ket_hotel, foto, created_at`

func scanHotel(rows pgx.Rows) (domain.Hotel, error) {
	h := domain.Hotel{}
	err := rows.Scan(
		&h.ID, &h.Kota, &h.NamaHotel, &h.TipeHotel, &h.HargaHotel,
		&h.AlamatHotel, &h.Latitude, &h.Longitude, &h.KetHotel, &h.Foto, &h.CreatedAt,
	)
	return h, err
}

func (r *hotelRepository) FindAll(ctx context.Context) ([]domain.Hotel, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+hotelColumns+` FROM hotel ORDER BY kota, tipe_hotel`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Hotel
	for rows.Next() {
		h, err := scanHotel(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, h)
	}
	return result, rows.Err()
}

func (r *hotelRepository) FindByID(ctx context.Context, id int) (*domain.Hotel, error) {
	h := &domain.Hotel{}
	err := r.db.QueryRow(ctx,
		`SELECT `+hotelColumns+` FROM hotel WHERE id = $1`, id,
	).Scan(
		&h.ID, &h.Kota, &h.NamaHotel, &h.TipeHotel, &h.HargaHotel,
		&h.AlamatHotel, &h.Latitude, &h.Longitude, &h.KetHotel, &h.Foto, &h.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return h, nil
}

func (r *hotelRepository) FindByKota(ctx context.Context, kota string) ([]domain.Hotel, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+hotelColumns+` FROM hotel WHERE kota = $1 ORDER BY tipe_hotel`, kota,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Hotel
	for rows.Next() {
		h, err := scanHotel(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, h)
	}
	return result, rows.Err()
}

func (r *hotelRepository) FindByKotaAndTipe(ctx context.Context, kota, tipe string) ([]domain.Hotel, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+hotelColumns+` FROM hotel WHERE kota = $1 AND tipe_hotel = $2 ORDER BY harga_hotel`,
		kota, tipe,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Hotel
	for rows.Next() {
		h, err := scanHotel(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, h)
	}
	return result, rows.Err()
}

func (r *hotelRepository) Create(ctx context.Context, h *domain.Hotel) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO hotel
			(kota, nama_hotel, tipe_hotel, harga_hotel, alamat_hotel, latitude, longitude, ket_hotel, foto)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at`,
		h.Kota, h.NamaHotel, h.TipeHotel, h.HargaHotel,
		h.AlamatHotel, h.Latitude, h.Longitude, h.KetHotel, h.Foto,
	).Scan(&h.ID, &h.CreatedAt)
}

func (r *hotelRepository) Update(ctx context.Context, h *domain.Hotel) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE hotel
		SET kota = $1, nama_hotel = $2, tipe_hotel = $3, harga_hotel = $4,
		    alamat_hotel = $5, latitude = $6, longitude = $7, ket_hotel = $8, foto = $9
		WHERE id = $10`,
		h.Kota, h.NamaHotel, h.TipeHotel, h.HargaHotel,
		h.AlamatHotel, h.Latitude, h.Longitude, h.KetHotel, h.Foto, h.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (r *hotelRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM hotel WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}