package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type WisataHandler struct {
	svc service.WisataService
}

func NewWisataHandler(svc service.WisataService) *WisataHandler {
	return &WisataHandler{svc}
}

// GET /api/wisata?kota=Tokyo&kategori=alam
func (h *WisataHandler) GetAll(c fiber.Ctx) error {
	kota := c.Query("kota")
	kategori := c.Query("kategori")

	if kota != "" && kategori != "" {
		data, err := h.svc.GetByKotaAndKategori(c.Context(), kota, kategori)
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

// POST /api/admin/wisata
func (h *WisataHandler) Create(c fiber.Ctx) error {
	var body domain.Wisata
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	if err := h.svc.Create(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.Created(c, "wisata berhasil ditambahkan", body)
}

// PUT /api/admin/wisata/:id
func (h *WisataHandler) Update(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var body domain.Wisata
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.ID = id

	if err := h.svc.Update(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "wisata berhasil diperbarui", body)
}

// DELETE /api/admin/wisata/:id
func (h *WisataHandler) Delete(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "wisata berhasil dihapus", nil)
}