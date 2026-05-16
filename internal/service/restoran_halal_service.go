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
	Create(ctx context.Context, r *domain.RestoranHalal) error
	Update(ctx context.Context, r *domain.RestoranHalal) error
	Delete(ctx context.Context, id int) error
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

func (s *restoranHalalService) Create(ctx context.Context, r *domain.RestoranHalal) error {
	return s.repo.Create(ctx, r)
}

func (s *restoranHalalService) Update(ctx context.Context, r *domain.RestoranHalal) error {
	return s.repo.Update(ctx, r)
}

func (s *restoranHalalService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}