package service

import (
	"context"

	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
)

type HotelService interface {
	GetAll(ctx context.Context) ([]domain.Hotel, error)
	GetByID(ctx context.Context, id int) (*domain.Hotel, error)
	GetByKota(ctx context.Context, kota string) ([]domain.Hotel, error)
	GetByKotaAndTipe(ctx context.Context, kota, tipe string) ([]domain.Hotel, error)
	Create(ctx context.Context, h *domain.Hotel) error
	Update(ctx context.Context, h *domain.Hotel) error
	Delete(ctx context.Context, id int) error
}

type hotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) HotelService {
	return &hotelService{repo}
}

func (s *hotelService) GetAll(ctx context.Context) ([]domain.Hotel, error) {
	return s.repo.FindAll(ctx)
}

func (s *hotelService) GetByID(ctx context.Context, id int) (*domain.Hotel, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *hotelService) GetByKota(ctx context.Context, kota string) ([]domain.Hotel, error) {
	return s.repo.FindByKota(ctx, kota)
}

func (s *hotelService) GetByKotaAndTipe(ctx context.Context, kota, tipe string) ([]domain.Hotel, error) {
	return s.repo.FindByKotaAndTipe(ctx, kota, tipe)
}

func (s *hotelService) Create(ctx context.Context, h *domain.Hotel) error {
	return s.repo.Create(ctx, h)
}

func (s *hotelService) Update(ctx context.Context, h *domain.Hotel) error {
	return s.repo.Update(ctx, h)
}

func (s *hotelService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}