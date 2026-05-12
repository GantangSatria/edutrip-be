package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type RestoranHalalService interface {
	GetAll(ctx context.Context) ([]domain.RestoranHalal, error)
	GetByID(ctx context.Context, id int) (*domain.RestoranHalal, error)
	GetByKota(ctx context.Context, kota string) ([]domain.RestoranHalal, error)
}

type restoranHalalService struct {
	repo repository.RestoranHalalRepository
}

func NewRestoranHalalService(repo repository.RestoranHalalRepository) RestoranHalalService {
	return &restoranHalalService{repo}
}

func (s *restoranHalalService) GetAll(ctx context.Context) ([]domain.RestoranHalal, error) {
	return s.repo.FindAll(ctx)
}

func (s *restoranHalalService) GetByID(ctx context.Context, id int) (*domain.RestoranHalal, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *restoranHalalService) GetByKota(ctx context.Context, kota string) ([]domain.RestoranHalal, error) {
	return s.repo.FindByKota(ctx, kota)
}