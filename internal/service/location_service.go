package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
	"github.com/GantangSatria/edutrip-be/pkg/dto/request"
)

type LocationService interface {
	GetAll(ctx context.Context) ([]domain.Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Location, error)
	Create(ctx context.Context, adminID uuid.UUID, req request.CreateLocation) (*domain.Location, error)
	Update(ctx context.Context, id uuid.UUID, req request.UpdateLocation) (*domain.Location, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo}
}

func (s *locationService) GetAll(ctx context.Context) ([]domain.Location, error) {
	return s.repo.FindAll(ctx)
}

func (s *locationService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Location, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *locationService) Create(ctx context.Context, adminID uuid.UUID, req request.CreateLocation) (*domain.Location, error) {
	loc := &domain.Location{
		AdminID:     adminID,
		Name:        req.Name,
		City:        req.City,
		Category:    req.Category,
		Address:     req.Address,
		Price:       req.Price,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Description: req.Description,
		Website:     req.Website,
		IsHalal:     req.IsHalal,
	}
	if err := s.repo.Create(ctx, loc); err != nil {
		return nil, err
	}
	return loc, nil
}

func (s *locationService) Update(ctx context.Context, id uuid.UUID, req request.UpdateLocation) (*domain.Location, error) {
	loc, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != ""        { loc.Name = req.Name }
	if req.City != ""        { loc.City = req.City }
	if req.Category != ""    { loc.Category = req.Category }
	if req.Address != ""     { loc.Address = req.Address }
	if req.Price > 0         { loc.Price = req.Price }
	if req.Latitude != 0     { loc.Latitude = req.Latitude }
	if req.Longitude != 0    { loc.Longitude = req.Longitude }
	if req.Description != "" { loc.Description = req.Description }
	if req.Website != ""     { loc.Website = req.Website }
	loc.IsHalal = req.IsHalal

	if err := s.repo.Update(ctx, loc); err != nil {
		return nil, err
	}
	return loc, nil
}

func (s *locationService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}