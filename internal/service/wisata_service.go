package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type WisataService interface {
	GetAll(ctx context.Context) ([]domain.Wisata, error)
	GetByID(ctx context.Context, id int) (*domain.Wisata, error)
	GetByKota(ctx context.Context, kota string) ([]domain.Wisata, error)
	GetByKotaAndKategori(ctx context.Context, kota, kategori string) ([]domain.Wisata, error)
	Create(ctx context.Context, w *domain.Wisata) error
	Update(ctx context.Context, w *domain.Wisata) error
	Delete(ctx context.Context, id int) error
}

type wisataService struct {
	repo repository.WisataRepository
}

func NewWisataService(repo repository.WisataRepository) WisataService {
	return &wisataService{repo}
}

func (s *wisataService) GetAll(ctx context.Context) ([]domain.Wisata, error) {
	return s.repo.FindAll(ctx)
}

func (s *wisataService) GetByID(ctx context.Context, id int) (*domain.Wisata, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *wisataService) GetByKota(ctx context.Context, kota string) ([]domain.Wisata, error) {
	return s.repo.FindByKota(ctx, kota)
}

func (s *wisataService) GetByKotaAndKategori(ctx context.Context, kota, kategori string) ([]domain.Wisata, error) {
	return s.repo.FindByKotaAndKategori(ctx, kota, kategori)
}

func (s *wisataService) Create(ctx context.Context, w *domain.Wisata) error {
	return s.repo.Create(ctx, w)
}

func (s *wisataService) Update(ctx context.Context, w *domain.Wisata) error {
	return s.repo.Update(ctx, w)
}

func (s *wisataService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}