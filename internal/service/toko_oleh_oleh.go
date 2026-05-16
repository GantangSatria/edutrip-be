package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type TokoOlehOlehService interface {
	GetAll(ctx context.Context) ([]domain.TokoOlehOleh, error)
	GetByID(ctx context.Context, id int) (*domain.TokoOlehOleh, error)
	GetByKota(ctx context.Context, kota string) ([]domain.TokoOlehOleh, error)
	Create(ctx context.Context, t *domain.TokoOlehOleh) error
	Update(ctx context.Context, t *domain.TokoOlehOleh) error
	Delete(ctx context.Context, id int) error
}

type tokoOlehOlehService struct {
	repo repository.TokoOlehOlehRepository
}

func NewTokoOlehOlehService(repo repository.TokoOlehOlehRepository) TokoOlehOlehService {
	return &tokoOlehOlehService{repo}
}

func (s *tokoOlehOlehService) GetAll(ctx context.Context) ([]domain.TokoOlehOleh, error) {
	return s.repo.FindAll(ctx)
}

func (s *tokoOlehOlehService) GetByID(ctx context.Context, id int) (*domain.TokoOlehOleh, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *tokoOlehOlehService) GetByKota(ctx context.Context, kota string) ([]domain.TokoOlehOleh, error) {
	return s.repo.FindByKota(ctx, kota)
}

func (s *tokoOlehOlehService) Create(ctx context.Context, t *domain.TokoOlehOleh) error {
	return s.repo.Create(ctx, t)
}

func (s *tokoOlehOlehService) Update(ctx context.Context, t *domain.TokoOlehOleh) error {
	return s.repo.Update(ctx, t)
}

func (s *tokoOlehOlehService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}