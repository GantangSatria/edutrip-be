package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
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

// POST /api/admin/fasilitas-ibadah
func (h *FasilitasIbadahHandler) Create(c fiber.Ctx) error {
	var body domain.FasilitasIbadah
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	if err := h.svc.Create(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.Created(c, "fasilitas ibadah berhasil ditambahkan", body)
}

// PUT /api/admin/fasilitas-ibadah/:id
func (h *FasilitasIbadahHandler) Update(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var body domain.FasilitasIbadah
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.ID = id

	if err := h.svc.Update(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "fasilitas ibadah berhasil diperbarui", body)
}

// DELETE /api/admin/fasilitas-ibadah/:id
func (h *FasilitasIbadahHandler) Delete(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "fasilitas ibadah berhasil dihapus", nil)
}