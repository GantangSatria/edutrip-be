package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type TransportasiRepository interface {
	FindAll(ctx context.Context) ([]domain.Transportasi, error)
	FindByID(ctx context.Context, id int) (*domain.Transportasi, error)
	FindByRute(ctx context.Context, rute string) ([]domain.Transportasi, error)
}

type transportasiRepository struct {
	db *pgxpool.Pool
}

func NewTransportasiRepository(db *pgxpool.Pool) TransportasiRepository {
	return &transportasiRepository{db}
}

const transportasiColumns = `id, jenis_transportasi, nama_transportasi, rute, kode_bandara,
	harga_transportasi_idr, ket_transportasi, foto, created_at`

func scanTransportasi(row pgx.Row) (*domain.Transportasi, error) {
	t := &domain.Transportasi{}
	err := row.Scan(
		&t.ID, &t.JenisTransportasi, &t.NamaTransportasi, &t.Rute, &t.KodeBandara,
		&t.HargaTransportasiIDR, &t.KetTransportasi, &t.Foto, &t.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *transportasiRepository) FindAll(ctx context.Context) ([]domain.Transportasi, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+transportasiColumns+` FROM transportasi ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Transportasi
	for rows.Next() {
		t := domain.Transportasi{}
		if err := rows.Scan(
			&t.ID, &t.JenisTransportasi, &t.NamaTransportasi, &t.Rute, &t.KodeBandara,
			&t.HargaTransportasiIDR, &t.KetTransportasi, &t.Foto, &t.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, rows.Err()
}

func (r *transportasiRepository) FindByID(ctx context.Context, id int) (*domain.Transportasi, error) {
	row := r.db.QueryRow(ctx,
		`SELECT `+transportasiColumns+` FROM transportasi WHERE id = $1`, id,
	)
	t, err := scanTransportasi(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return t, nil
}

func (r *transportasiRepository) FindByRute(ctx context.Context, rute string) ([]domain.Transportasi, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+transportasiColumns+` FROM transportasi WHERE rute ILIKE $1 ORDER BY harga_transportasi_idr`,
		"%"+rute+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Transportasi
	for rows.Next() {
		t := domain.Transportasi{}
		if err := rows.Scan(
			&t.ID, &t.JenisTransportasi, &t.NamaTransportasi, &t.Rute, &t.KodeBandara,
			&t.HargaTransportasiIDR, &t.KetTransportasi, &t.Foto, &t.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, rows.Err()
}