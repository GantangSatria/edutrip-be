package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/request"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
)

type TripHandler struct {
	svc service.TripService
}

func NewTripHandler(svc service.TripService) *TripHandler {
	return &TripHandler{svc}
}

func (h *TripHandler) GetAll(c fiber.Ctx) error {
	trips, err := h.svc.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "failed to fetch trips")
	}
	return response.OK(c, "trips retrieved", trips)
}

func (h *TripHandler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid trip id")
	}

	trip, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "trip not found")
	}
	return response.OK(c, "trip retrieved", trip)
}

func (h *TripHandler) Create(c fiber.Ctx) error {
	var req request.CreateTrip
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	adminID, err := uuid.Parse(c.Locals("adminID").(string))
	if err != nil {
		return response.Unauthorized(c, "invalid admin token")
	}

	trip, err := h.svc.Create(c.Context(), adminID, req)
	if err != nil {
		return response.InternalError(c, "failed to create trip")
	}
	return response.Created(c, "trip created", trip)
}

func (h *TripHandler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid trip id")
	}

	var req request.UpdateTrip
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	trip, err := h.svc.Update(c.Context(), id, req)
	if err != nil {
		return response.NotFound(c, "trip not found")
	}
	return response.OK(c, "trip updated", trip)
}

func (h *TripHandler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid trip id")
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return response.NotFound(c, "trip not found")
	}
	return response.OK(c, "trip deleted", nil)
}