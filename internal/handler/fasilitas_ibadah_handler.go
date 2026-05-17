// ============================================================
// fasilitas_ibadah_handler.go
// ============================================================
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

// GetAll godoc
// @Summary      List fasilitas ibadah
// @Description  Ambil semua fasilitas ibadah. Filter opsional dengan query param kota.
// @Tags         Fasilitas Ibadah
// @Produce      json
// @Param        kota  query     string  false  "Filter by kota"
// @Success      200   {object}  response.Base{data=[]domain.FasilitasIbadah}
// @Failure      500   {object}  response.Base
// @Router       /fasilitas-ibadah [get]
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

// GetByID godoc
// @Summary      Detail fasilitas ibadah
// @Description  Ambil satu fasilitas ibadah berdasarkan ID
// @Tags         Fasilitas Ibadah
// @Produce      json
// @Param        id   path      int  true  "Fasilitas Ibadah ID"
// @Success      200  {object}  response.Base{data=domain.FasilitasIbadah}
// @Failure      400  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /fasilitas-ibadah/{id} [get]
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

// Create godoc
// @Summary      Tambah fasilitas ibadah
// @Description  Tambah data fasilitas ibadah baru (admin only)
// @Tags         Fasilitas Ibadah
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      domain.FasilitasIbadah  true  "Data fasilitas ibadah"
// @Success      201   {object}  response.Base{data=domain.FasilitasIbadah}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Router       /admin/fasilitas-ibadah [post]
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

// Update godoc
// @Summary      Update fasilitas ibadah
// @Description  Perbarui data fasilitas ibadah berdasarkan ID (admin only)
// @Tags         Fasilitas Ibadah
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                     true  "Fasilitas Ibadah ID"
// @Param        body  body      domain.FasilitasIbadah  true  "Data fasilitas ibadah"
// @Success      200   {object}  response.Base{data=domain.FasilitasIbadah}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/fasilitas-ibadah/{id} [put]
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

// Delete godoc
// @Summary      Hapus fasilitas ibadah
// @Description  Hapus data fasilitas ibadah berdasarkan ID (admin only)
// @Tags         Fasilitas Ibadah
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Fasilitas Ibadah ID"
// @Success      200  {object}  response.Base
// @Failure      400  {object}  response.Base
// @Failure      401  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /admin/fasilitas-ibadah/{id} [delete]
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