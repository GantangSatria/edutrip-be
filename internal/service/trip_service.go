package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/repository"
	"github.com/GantangSatria/edutrip-be/pkg/dto/request"
)

type TripService interface {
	GetAll(ctx context.Context) ([]domain.Trip, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Trip, error)
	Create(ctx context.Context, adminID uuid.UUID, req request.CreateTrip) (*domain.Trip, error)
	Update(ctx context.Context, id uuid.UUID, req request.UpdateTrip) (*domain.Trip, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type tripService struct {
	repo repository.TripRepository
}

func NewTripService(repo repository.TripRepository) TripService {
	return &tripService{repo}
}

func (s *tripService) GetAll(ctx context.Context) ([]domain.Trip, error) {
	return s.repo.FindAll(ctx)
}

func (s *tripService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *tripService) Create(ctx context.Context, adminID uuid.UUID, req request.CreateTrip) (*domain.Trip, error) {
	status := req.Status
	if status == "" {
		status = "draft"
	}
	trip := &domain.Trip{
		AdminID:          adminID,
		Title:            req.Title,
		Description:      req.Description,
		City:             req.City,
		Theme:            req.Theme,
		DurationDays:     req.DurationDays,
		Price:            req.Price,
		Facilities:       req.Facilities,
		ItinerarySummary: req.ItinerarySummary,
		Status:           status,
	}
	if err := s.repo.Create(ctx, trip); err != nil {
		return nil, err
	}
	return trip, nil
}

func (s *tripService) Update(ctx context.Context, id uuid.UUID, req request.UpdateTrip) (*domain.Trip, error) {
	trip, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != ""            { trip.Title = req.Title }
	if req.Description != ""      { trip.Description = req.Description }
	if req.City != ""             { trip.City = req.City }
	if req.Theme != ""            { trip.Theme = req.Theme }
	if req.DurationDays > 0       { trip.DurationDays = req.DurationDays }
	if req.Price > 0              { trip.Price = req.Price }
	if req.Facilities != ""       { trip.Facilities = req.Facilities }
	if req.ItinerarySummary != "" { trip.ItinerarySummary = req.ItinerarySummary }
	if req.Status != ""           { trip.Status = req.Status }

	if err := s.repo.Update(ctx, trip); err != nil {
		return nil, err
	}
	return trip, nil
}

func (s *tripService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}