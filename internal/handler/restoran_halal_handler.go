package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/gofiber/fiber/v3"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type RestoranHalalHandler struct {
	svc service.RestoranHalalService
}

func NewRestoranHalalHandler(svc service.RestoranHalalService) *RestoranHalalHandler {
	return &RestoranHalalHandler{svc}
}

// GET /api/restoran-halal?kota=Tokyo
func (h *RestoranHalalHandler) GetAll(c fiber.Ctx) error {
	kota := c.Query("kota")

	if kota != "" {
		data, err := h.svc.GetByKota(c.Context(), kota)
		if err != nil {
			return utils.HandleServiceError(c, err)
		}
		return response.OK(c, "success", data)
	}

	data, err := h.svc.GetAll(c.Context())
	if err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "success", data)
}

// GET /api/restoran-halal/:id
func (h *RestoranHalalHandler) GetByID(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	data, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "success", data)
}