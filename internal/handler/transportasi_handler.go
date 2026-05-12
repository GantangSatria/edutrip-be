package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/gofiber/fiber/v3"
)

type TransportasiHandler struct {
	svc service.TransportasiService
}

func NewTransportasiHandler(svc service.TransportasiService) *TransportasiHandler {
	return &TransportasiHandler{svc}
}

// GET /api/transportasi
func (h *TransportasiHandler) GetAll(c fiber.Ctx) error {
	rute := c.Query("rute")

	if rute != "" {
		data, err := h.svc.GetByRute(c.Context(), rute)
		if err != nil {
			return handleServiceError(c, err)
		}
		return response.OK(c, "success", data)
	}

	data, err := h.svc.GetAll(c.Context())
	if err != nil {
		return handleServiceError(c, err)
	}
	return response.OK(c, "success", data)
}

// GET /api/transportasi/:id
func (h *TransportasiHandler) GetByID(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	data, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return handleServiceError(c, err)
	}
	return response.OK(c, "success", data)
}