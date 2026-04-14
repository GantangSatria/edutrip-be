package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/request"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
)

type LocationHandler struct {
	svc service.LocationService
}

func NewLocationHandler(svc service.LocationService) *LocationHandler {
	return &LocationHandler{svc}
}

func (h *LocationHandler) GetAll(c fiber.Ctx) error {
	locs, err := h.svc.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "failed to fetch locations")
	}
	return response.OK(c, "locations retrieved", locs)
}

func (h *LocationHandler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid location id")
	}

	loc, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "location not found")
	}
	return response.OK(c, "location retrieved", loc)
}

func (h *LocationHandler) Create(c fiber.Ctx) error {
	var req request.CreateLocation
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	adminID, err := uuid.Parse(c.Locals("adminID").(string))
	if err != nil {
		return response.Unauthorized(c, "invalid admin token")
	}

	loc, err := h.svc.Create(c.Context(), adminID, req)
	if err != nil {
		return response.InternalError(c, "failed to create location")
	}
	return response.Created(c, "location created", loc)
}

func (h *LocationHandler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid location id")
	}

	var req request.UpdateLocation
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	loc, err := h.svc.Update(c.Context(), id, req)
	if err != nil {
		return response.NotFound(c, "location not found")
	}
	return response.OK(c, "location updated", loc)
}

func (h *LocationHandler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid location id")
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return response.NotFound(c, "location not found")
	}
	return response.OK(c, "location deleted", nil)
}