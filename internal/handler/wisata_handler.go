package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/gofiber/fiber/v3"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
)

type WisataHandler struct {
	svc service.WisataService
}

func NewWisataHandler(svc service.WisataService) *WisataHandler {
	return &WisataHandler{svc}
}

// GET /api/wisata?kota=Tokyo&kategori=Museum
func (h *WisataHandler) GetAll(c fiber.Ctx) error {
	kota := c.Query("kota")
	kategori := c.Query("kategori")

	switch {
	case kota != "" && kategori != "":
		data, err := h.svc.GetByKotaAndKategori(c.Context(), kota, kategori)
		if err != nil {
			return utils.HandleServiceError(c, err)
		}
		return response.OK(c, "success", data)
	case kota != "":
		data, err := h.svc.GetByKota(c.Context(), kota)
		if err != nil {
			return utils.HandleServiceError(c, err)
		}
		return response.OK(c, "success", data)
	default:
		data, err := h.svc.GetAll(c.Context())
		if err != nil {
			return utils.HandleServiceError(c, err)
		}
		return response.OK(c, "success", data)
	}
}

// GET /api/wisata/:id
func (h *WisataHandler) GetByID(c fiber.Ctx) error {
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