package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SettingRepository interface {
	FindAll(ctx context.Context) ([]domain.Setting, error)
	FindByKey(ctx context.Context, key string) (*domain.Setting, error)
	Update(ctx context.Context, s *domain.Setting) error
}

type settingRepository struct {
	db *pgxpool.Pool
}

func NewSettingRepository(db *pgxpool.Pool) SettingRepository {
	return &settingRepository{db}
}

const settingColumns = `key, value, updated_at`

func scanSetting(rows pgx.Rows) (domain.Setting, error) {
	s := domain.Setting{}
	err := rows.Scan(&s.Key, &s.Value, &s.UpdatedAt)
	return s, err
}

func (r *settingRepository) FindAll(ctx context.Context) ([]domain.Setting, error) {
	rows, err := r.db.Query(ctx, `SELECT `+settingColumns+` FROM settings ORDER BY key`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Setting
	for rows.Next() {
		s, err := scanSetting(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, rows.Err()
}

func (r *settingRepository) FindByKey(ctx context.Context, key string) (*domain.Setting, error) {
	s := &domain.Setting{}
	err := r.db.QueryRow(ctx, `SELECT `+settingColumns+` FROM settings WHERE key = $1`, key).
		Scan(&s.Key, &s.Value, &s.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return s, nil
}

func (r *settingRepository) Update(ctx context.Context, s *domain.Setting) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE settings SET value = $1, updated_at = NOW() WHERE key = $2`,
		s.Value, s.Key,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}
