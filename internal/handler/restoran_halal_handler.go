package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type RestoranHalalHandler struct {
	svc service.RestoranHalalService
}

func NewRestoranHalalHandler(svc service.RestoranHalalService) *RestoranHalalHandler {
	return &RestoranHalalHandler{svc}
}

// GetAll godoc
// @Summary      List restoran halal
// @Description  Ambil semua restoran halal. Filter opsional dengan query param kota.
// @Tags         Restoran Halal
// @Produce      json
// @Param        kota  query     string  false  "Filter by kota"
// @Success      200   {object}  response.Base{data=[]domain.RestoranHalal}
// @Failure      500   {object}  response.Base
// @Router       /restoran-halal [get]
func (h *RestoranHalalHandler) GetAll(c fiber.Ctx) error {
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
// @Summary      Detail restoran halal
// @Description  Ambil satu restoran halal berdasarkan ID
// @Tags         Restoran Halal
// @Produce      json
// @Param        id   path      int  true  "Restoran Halal ID"
// @Success      200  {object}  response.Base{data=domain.RestoranHalal}
// @Failure      400  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /restoran-halal/{id} [get]
func (h *RestoranHalalHandler) GetByID(c fiber.Ctx) error {
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
// @Summary      Tambah restoran halal
// @Description  Tambah data restoran halal baru (admin only)
// @Tags         Restoran Halal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      domain.RestoranHalal  true  "Data restoran halal"
// @Success      201   {object}  response.Base{data=domain.RestoranHalal}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Router       /admin/restoran-halal [post]
func (h *RestoranHalalHandler) Create(c fiber.Ctx) error {
	var body domain.RestoranHalal
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	if err := h.svc.Create(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.Created(c, "restoran halal berhasil ditambahkan", body)
}

// Update godoc
// @Summary      Update restoran halal
// @Description  Perbarui data restoran halal berdasarkan ID (admin only)
// @Tags         Restoran Halal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                   true  "Restoran Halal ID"
// @Param        body  body      domain.RestoranHalal  true  "Data restoran halal"
// @Success      200   {object}  response.Base{data=domain.RestoranHalal}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/restoran-halal/{id} [put]
func (h *RestoranHalalHandler) Update(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var body domain.RestoranHalal
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.ID = id

	if err := h.svc.Update(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "restoran halal berhasil diperbarui", body)
}

// Delete godoc
// @Summary      Hapus restoran halal
// @Description  Hapus data restoran halal berdasarkan ID (admin only)
// @Tags         Restoran Halal
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Restoran Halal ID"
// @Success      200  {object}  response.Base
// @Failure      400  {object}  response.Base
// @Failure      401  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /admin/restoran-halal/{id} [delete]
func (h *RestoranHalalHandler) Delete(c fiber.Ctx) error {
	id, err := utils.ParseID(c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "restoran halal berhasil dihapus", nil)
}