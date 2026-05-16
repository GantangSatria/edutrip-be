package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/gofiber/fiber/v3"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type FasilitasIbadahHandler struct {
	svc service.FasilitasIbadahService
}

func NewFasilitasIbadahHandler(svc service.FasilitasIbadahService) *FasilitasIbadahHandler {
	return &FasilitasIbadahHandler{svc}
}

// GET /api/fasilitas-ibadah?kota=Tokyo
func (h *FasilitasIbadahHandler) GetAll(c fiber.Ctx) error {
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

// GET /api/fasilitas-ibadah/:id
func (h *FasilitasIbadahHandler) GetByID(c fiber.Ctx) error {
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