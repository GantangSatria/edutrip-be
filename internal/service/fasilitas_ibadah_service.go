package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type FasilitasIbadahService interface {
	GetAll(ctx context.Context) ([]domain.FasilitasIbadah, error)
	GetByID(ctx context.Context, id int) (*domain.FasilitasIbadah, error)
	GetByKota(ctx context.Context, kota string) ([]domain.FasilitasIbadah, error)
	Create(ctx context.Context, f *domain.FasilitasIbadah) error
	Update(ctx context.Context, f *domain.FasilitasIbadah) error
	Delete(ctx context.Context, id int) error
}

type fasilitasIbadahService struct {
	repo repository.FasilitasIbadahRepository
}

func NewFasilitasIbadahService(repo repository.FasilitasIbadahRepository) FasilitasIbadahService {
	return &fasilitasIbadahService{repo}
}

func (s *fasilitasIbadahService) GetAll(ctx context.Context) ([]domain.FasilitasIbadah, error) {
	return s.repo.FindAll(ctx)
}

func (s *fasilitasIbadahService) GetByID(ctx context.Context, id int) (*domain.FasilitasIbadah, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *fasilitasIbadahService) GetByKota(ctx context.Context, kota string) ([]domain.FasilitasIbadah, error) {
	return s.repo.FindByKota(ctx, kota)
}

func (s *fasilitasIbadahService) Create(ctx context.Context, f *domain.FasilitasIbadah) error {
	return s.repo.Create(ctx, f)
}

func (s *fasilitasIbadahService) Update(ctx context.Context, f *domain.FasilitasIbadah) error {
	return s.repo.Update(ctx, f)
}

func (s *fasilitasIbadahService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}