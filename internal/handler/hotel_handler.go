package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type HotelHandler struct {
	svc service.HotelService
}

func NewHotelHandler(svc service.HotelService) *HotelHandler {
	return &HotelHandler{svc}
}

// GET /api/hotel?kota=Tokyo&tipe=bintang3
func (h *HotelHandler) GetAll(c fiber.Ctx) error {
	kota := c.Query("kota")
	tipe := c.Query("tipe")

	if kota != "" && tipe != "" {
		data, err := h.svc.GetByKotaAndTipe(c.Context(), kota, tipe)
		if err != nil {
			return utils.HandleServiceError(c, err)
		}
		return response.OK(c, "success", data)
	}
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

// GET /api/hotel/:id
func (h *HotelHandler) GetByID(c fiber.Ctx) error {
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

// POST /api/admin/hotel
func (h *HotelHandler) Create(c fiber.Ctx) error {
	var body domain.Hotel
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	if err := h.svc.Create(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.Created(c, "hotel berhasil ditambahkan", body)
}

// PUT /api/admin/hotel/:id
func (h *HotelHandler) Update(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var body domain.Hotel
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.ID = id

	if err := h.svc.Update(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "hotel berhasil diperbarui", body)
}

// DELETE /api/admin/hotel/:id
func (h *HotelHandler) Delete(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "hotel berhasil dihapus", nil)
}