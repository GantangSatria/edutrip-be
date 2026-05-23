package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type KotaService interface {
	GetAllKota(ctx context.Context) ([]domain.Kota, error)
	GetKotaByID(ctx context.Context, id int) (*domain.Kota, error)
	CreateKota(ctx context.Context, k *domain.Kota) error
	UpdateKota(ctx context.Context, k *domain.Kota) error
	DeleteKota(ctx context.Context, id int) error
}

type kotaService struct {
	repo repository.KotaRepository
}

func NewKotaService(repo repository.KotaRepository) KotaService {
	return &kotaService{repo}
}

func (s *kotaService) GetAllKota(ctx context.Context) ([]domain.Kota, error) {
	return s.repo.FindAll(ctx)
}

func (s *kotaService) GetKotaByID(ctx context.Context, id int) (*domain.Kota, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *kotaService) CreateKota(ctx context.Context, k *domain.Kota) error {
	return s.repo.Create(ctx, k)
}

func (s *kotaService) UpdateKota(ctx context.Context, k *domain.Kota) error {
	return s.repo.Update(ctx, k)
}

func (s *kotaService) DeleteKota(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
