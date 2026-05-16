package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
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

// POST /api/admin/toko-oleh-oleh
func (h *TokoOlehOlehHandler) Create(c fiber.Ctx) error {
	var body domain.TokoOlehOleh
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	if err := h.svc.Create(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.Created(c, "toko oleh-oleh berhasil ditambahkan", body)
}

// PUT /api/admin/toko-oleh-oleh/:id
func (h *TokoOlehOlehHandler) Update(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var body domain.TokoOlehOleh
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.ID = id

	if err := h.svc.Update(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "toko oleh-oleh berhasil diperbarui", body)
}

// DELETE /api/admin/toko-oleh-oleh/:id
func (h *TokoOlehOlehHandler) Delete(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "toko oleh-oleh berhasil dihapus", nil)
}