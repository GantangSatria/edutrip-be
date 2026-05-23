package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type KotaHandler struct {
	svc service.KotaService
}

func NewKotaHandler(svc service.KotaService) *KotaHandler {
	return &KotaHandler{svc}
}

// GetAll godoc
// @Summary      List kota
// @Description  Ambil semua kota
// @Tags         Kota
// @Produce      json
// @Success      200  {object}  response.Base{data=[]domain.Kota}
// @Failure      500  {object}  response.Base
// @Router       /kota [get]
func (h *KotaHandler) GetAll(c fiber.Ctx) error {
	data, err := h.svc.GetAllKota(c.Context())
	if err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "success", data)
}

// GetByID godoc
// @Summary      Detail kota
// @Description  Ambil satu kota berdasarkan ID
// @Tags         Kota
// @Produce      json
// @Param        id   path      int  true  "Kota ID"
// @Success      200  {object}  response.Base{data=domain.Kota}
// @Failure      400  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /kota/{id} [get]
func (h *KotaHandler) GetByID(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	data, err := h.svc.GetKotaByID(c.Context(), id)
	if err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "success", data)
}

// Create godoc
// @Summary      Tambah kota
// @Description  Tambah data kota baru (admin only)
// @Tags         Kota
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      domain.Kota  true  "Data kota"
// @Success      201   {object}  response.Base{data=domain.Kota}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Router       /admin/kota [post]
func (h *KotaHandler) Create(c fiber.Ctx) error {
	var body domain.Kota
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	if err := h.svc.CreateKota(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.Created(c, "kota berhasil ditambahkan", body)
}

// Update godoc
// @Summary      Update kota
// @Description  Perbarui data kota berdasarkan ID (admin only)
// @Tags         Kota
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int          true  "Kota ID"
// @Param        body  body      domain.Kota  true  "Data kota"
// @Success      200   {object}  response.Base{data=domain.Kota}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/kota/{id} [put]
func (h *KotaHandler) Update(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var body domain.Kota
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.ID = id

	if err := h.svc.UpdateKota(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "kota berhasil diperbarui", body)
}

// Delete godoc
// @Summary      Hapus kota
// @Description  Hapus data kota berdasarkan ID (admin only)
// @Tags         Kota
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Kota ID"
// @Success      200  {object}  response.Base
// @Failure      400  {object}  response.Base
// @Failure      401  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /admin/kota/{id} [delete]
func (h *KotaHandler) Delete(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.svc.DeleteKota(c.Context(), id); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "kota berhasil dihapus", nil)
}
