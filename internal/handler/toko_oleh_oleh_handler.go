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

// GetAll godoc
// @Summary      List toko oleh-oleh
// @Description  Ambil semua toko oleh-oleh. Filter opsional dengan query param kota.
// @Tags         Toko Oleh-Oleh
// @Produce      json
// @Param        kota  query     string  false  "Filter by kota"
// @Success      200   {object}  response.Base{data=[]domain.TokoOlehOleh}
// @Failure      500   {object}  response.Base
// @Router       /toko-oleh-oleh [get]
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

// GetByID godoc
// @Summary      Detail toko oleh-oleh
// @Description  Ambil satu toko oleh-oleh berdasarkan ID
// @Tags         Toko Oleh-Oleh
// @Produce      json
// @Param        id   path      int  true  "Toko Oleh-Oleh ID"
// @Success      200  {object}  response.Base{data=domain.TokoOlehOleh}
// @Failure      400  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /toko-oleh-oleh/{id} [get]
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

// Create godoc
// @Summary      Tambah toko oleh-oleh
// @Description  Tambah data toko oleh-oleh baru (admin only)
// @Tags         Toko Oleh-Oleh
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      domain.TokoOlehOleh  true  "Data toko oleh-oleh"
// @Success      201   {object}  response.Base{data=domain.TokoOlehOleh}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Router       /admin/toko-oleh-oleh [post]
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

// Update godoc
// @Summary      Update toko oleh-oleh
// @Description  Perbarui data toko oleh-oleh berdasarkan ID (admin only)
// @Tags         Toko Oleh-Oleh
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                  true  "Toko Oleh-Oleh ID"
// @Param        body  body      domain.TokoOlehOleh  true  "Data toko oleh-oleh"
// @Success      200   {object}  response.Base{data=domain.TokoOlehOleh}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/toko-oleh-oleh/{id} [put]
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

// Delete godoc
// @Summary      Hapus toko oleh-oleh
// @Description  Hapus data toko oleh-oleh berdasarkan ID (admin only)
// @Tags         Toko Oleh-Oleh
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Toko Oleh-Oleh ID"
// @Success      200  {object}  response.Base
// @Failure      400  {object}  response.Base
// @Failure      401  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /admin/toko-oleh-oleh/{id} [delete]
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