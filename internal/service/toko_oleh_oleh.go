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