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
	Create(ctx context.Context, t *domain.Transportasi) error
	Update(ctx context.Context, t *domain.Transportasi) error
	Delete(ctx context.Context, id int) error
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

func (r *transportasiRepository) Create(ctx context.Context, t *domain.Transportasi) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO transportasi
			(jenis_transportasi, nama_transportasi, rute, kode_bandara, harga_transportasi_idr, ket_transportasi, foto)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at`,
		t.JenisTransportasi, t.NamaTransportasi, t.Rute, t.KodeBandara,
		t.HargaTransportasiIDR, t.KetTransportasi, t.Foto,
	).Scan(&t.ID, &t.CreatedAt)
}

func (r *transportasiRepository) Update(ctx context.Context, t *domain.Transportasi) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE transportasi
		SET jenis_transportasi = $1, nama_transportasi = $2, rute = $3, kode_bandara = $4,
		    harga_transportasi_idr = $5, ket_transportasi = $6, foto = $7
		WHERE id = $8`,
		t.JenisTransportasi, t.NamaTransportasi, t.Rute, t.KodeBandara,
		t.HargaTransportasiIDR, t.KetTransportasi, t.Foto, t.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (r *transportasiRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM transportasi WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}