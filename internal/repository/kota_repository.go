package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type KotaRepository interface {
	FindAll(ctx context.Context) ([]domain.Kota, error)
	FindByID(ctx context.Context, id int) (*domain.Kota, error)
	Create(ctx context.Context, k *domain.Kota) error
	Update(ctx context.Context, k *domain.Kota) error
	Delete(ctx context.Context, id int) error
}

type kotaRepository struct {
	db *pgxpool.Pool
}

func NewKotaRepository(db *pgxpool.Pool) KotaRepository {
	return &kotaRepository{db}
}

const kotaColumns = `id, name, image, description, halal_spots_value, halal_spots_label, main_mosque_title, main_mosque_subtitle, features, terrain_lead, terrain_rest, created_at`

func scanKota(rows pgx.Rows) (domain.Kota, error) {
	k := domain.Kota{}
	err := rows.Scan(
		&k.ID, &k.Name, &k.Image, &k.Description, &k.HalalSpotsValue,
		&k.HalalSpotsLabel, &k.MainMosqueTitle, &k.MainMosqueSubtitle, &k.Features,
		&k.TerrainLead, &k.TerrainRest, &k.CreatedAt,
	)
	return k, err
}

func (r *kotaRepository) FindAll(ctx context.Context) ([]domain.Kota, error) {
	rows, err := r.db.Query(ctx, `SELECT `+kotaColumns+` FROM kota ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Kota
	for rows.Next() {
		k, err := scanKota(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, k)
	}
	return result, rows.Err()
}

func (r *kotaRepository) FindByID(ctx context.Context, id int) (*domain.Kota, error) {
	k := &domain.Kota{}
	err := r.db.QueryRow(ctx, `SELECT `+kotaColumns+` FROM kota WHERE id = $1`, id).
		Scan(
			&k.ID, &k.Name, &k.Image, &k.Description, &k.HalalSpotsValue,
			&k.HalalSpotsLabel, &k.MainMosqueTitle, &k.MainMosqueSubtitle, &k.Features,
			&k.TerrainLead, &k.TerrainRest, &k.CreatedAt,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return k, nil
}

func (r *kotaRepository) Create(ctx context.Context, k *domain.Kota) error {
	return r.db.QueryRow(ctx,
		`INSERT INTO kota
			(name, image, description, halal_spots_value, halal_spots_label, main_mosque_title, main_mosque_subtitle, features, terrain_lead, terrain_rest, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())
		RETURNING id, created_at`,
		k.Name, k.Image, k.Description, k.HalalSpotsValue, k.HalalSpotsLabel,
		k.MainMosqueTitle, k.MainMosqueSubtitle, k.Features, k.TerrainLead, k.TerrainRest,
	).Scan(&k.ID, &k.CreatedAt)
}

func (r *kotaRepository) Update(ctx context.Context, k *domain.Kota) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE kota
		SET name = $1, image = $2, description = $3, halal_spots_value = $4,
		    halal_spots_label = $5, main_mosque_title = $6, main_mosque_subtitle = $7,
		    features = $8, terrain_lead = $9, terrain_rest = $10
		WHERE id = $11`,
		k.Name, k.Image, k.Description, k.HalalSpotsValue,
		k.HalalSpotsLabel, k.MainMosqueTitle, k.MainMosqueSubtitle,
		k.Features, k.TerrainLead, k.TerrainRest, k.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (r *kotaRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM kota WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}
