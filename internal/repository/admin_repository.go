package repository

import (
	"context"
	"errors"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.Admin, error)
}

type adminRepository struct {
	db *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindByEmail(ctx context.Context, email string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	err := r.db.QueryRow(ctx,
		`SELECT id, name, email, password_hash, created_at FROM admins WHERE email = $1`,
		email,
	).Scan(&admin.ID, &admin.Name, &admin.Email, &admin.PasswordHash, &admin.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return admin, nil
}