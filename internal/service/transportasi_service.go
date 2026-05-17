package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type TransportasiService interface {
	GetAll(ctx context.Context) ([]domain.Transportasi, error)
	GetByID(ctx context.Context, id int) (*domain.Transportasi, error)
	GetByRute(ctx context.Context, rute string) ([]domain.Transportasi, error)
	Create(ctx context.Context, t *domain.Transportasi) error
	Update(ctx context.Context, t *domain.Transportasi) error
	Delete(ctx context.Context, id int) error
}

type transportasiService struct {
	repo repository.TransportasiRepository
}

func NewTransportasiService(repo repository.TransportasiRepository) TransportasiService {
	return &transportasiService{repo}
}

func (s *transportasiService) GetAll(ctx context.Context) ([]domain.Transportasi, error) {
	return s.repo.FindAll(ctx)
}

func (s *transportasiService) GetByID(ctx context.Context, id int) (*domain.Transportasi, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *transportasiService) GetByRute(ctx context.Context, rute string) ([]domain.Transportasi, error) {
	return s.repo.FindByRute(ctx, rute)
}

func (s *transportasiService) Create(ctx context.Context, t *domain.Transportasi) error {
	return s.repo.Create(ctx, t)
}

func (s *transportasiService) Update(ctx context.Context, t *domain.Transportasi) error {
	return s.repo.Update(ctx, t)
}

func (s *transportasiService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}