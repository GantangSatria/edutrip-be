package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type TransportasiHandler struct {
	svc service.TransportasiService
}

func NewTransportasiHandler(svc service.TransportasiService) *TransportasiHandler {
	return &TransportasiHandler{svc}
}

// GetAll godoc
// @Summary      List transportasi
// @Description  Ambil semua transportasi. Filter opsional dengan query param rute (ILIKE).
// @Tags         Transportasi
// @Produce      json
// @Param        rute  query     string  false  "Filter by rute (partial match)"
// @Success      200   {object}  response.Base{data=[]domain.Transportasi}
// @Failure      500   {object}  response.Base
// @Router       /transportasi [get]
func (h *TransportasiHandler) GetAll(c fiber.Ctx) error {
	rute := c.Query("rute")
	if rute != "" {
		data, err := h.svc.GetByRute(c.Context(), rute)
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
// @Summary      Detail transportasi
// @Description  Ambil satu transportasi berdasarkan ID
// @Tags         Transportasi
// @Produce      json
// @Param        id   path      int  true  "Transportasi ID"
// @Success      200  {object}  response.Base{data=domain.Transportasi}
// @Failure      400  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /transportasi/{id} [get]
func (h *TransportasiHandler) GetByID(c fiber.Ctx) error {
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
// @Summary      Tambah transportasi
// @Description  Tambah data transportasi baru (admin only)
// @Tags         Transportasi
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      domain.Transportasi  true  "Data transportasi"
// @Success      201   {object}  response.Base{data=domain.Transportasi}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Router       /admin/transportasi [post]
func (h *TransportasiHandler) Create(c fiber.Ctx) error {
	var body domain.Transportasi
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	if err := h.svc.Create(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.Created(c, "transportasi berhasil ditambahkan", body)
}

// Update godoc
// @Summary      Update transportasi
// @Description  Perbarui data transportasi berdasarkan ID (admin only)
// @Tags         Transportasi
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                  true  "Transportasi ID"
// @Param        body  body      domain.Transportasi  true  "Data transportasi"
// @Success      200   {object}  response.Base{data=domain.Transportasi}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/transportasi/{id} [put]
func (h *TransportasiHandler) Update(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var body domain.Transportasi
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.ID = id

	if err := h.svc.Update(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "transportasi berhasil diperbarui", body)
}

// Delete godoc
// @Summary      Hapus transportasi
// @Description  Hapus data transportasi berdasarkan ID (admin only)
// @Tags         Transportasi
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Transportasi ID"
// @Success      200  {object}  response.Base
// @Failure      400  {object}  response.Base
// @Failure      401  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /admin/transportasi/{id} [delete]
func (h *TransportasiHandler) Delete(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "transportasi berhasil dihapus", nil)
}