package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/gofiber/fiber/v3"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type TokoOlehOlehHandler struct {
	svc service.TokoOlehOlehService
}

func NewTokoOlehOlehHandler(svc service.TokoOlehOlehService) *TokoOlehOlehHandler {
	return &TokoOlehOlehHandler{svc}
}

// GET /api/toko-oleh-oleh?kota=Tokyo
func (h *TokoOlehOlehHandler) GetAll(c fiber.Ctx) error {
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

// GET /api/toko-oleh-oleh/:id
func (h *TokoOlehOlehHandler) GetByID(c fiber.Ctx) error {
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