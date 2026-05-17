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

// GetAll godoc
// @Summary      List wisata
// @Description  Ambil semua wisata. Filter opsional dengan query param kota dan/atau kategori.
// @Tags         Wisata
// @Produce      json
// @Param        kota      query     string  false  "Filter by kota"
// @Param        kategori  query     string  false  "Filter by kategori wisata"
// @Success      200       {object}  response.Base{data=[]domain.Wisata}
// @Failure      500       {object}  response.Base
// @Router       /wisata [get]
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

// GetByID godoc
// @Summary      Detail wisata
// @Description  Ambil satu wisata berdasarkan ID
// @Tags         Wisata
// @Produce      json
// @Param        id   path      int  true  "Wisata ID"
// @Success      200  {object}  response.Base{data=domain.Wisata}
// @Failure      400  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /wisata/{id} [get]
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

// Create godoc
// @Summary      Tambah wisata
// @Description  Tambah data wisata baru (admin only)
// @Tags         Wisata
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      domain.Wisata  true  "Data wisata"
// @Success      201   {object}  response.Base{data=domain.Wisata}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Router       /admin/wisata [post]
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

// Update godoc
// @Summary      Update wisata
// @Description  Perbarui data wisata berdasarkan ID (admin only)
// @Tags         Wisata
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int            true  "Wisata ID"
// @Param        body  body      domain.Wisata  true  "Data wisata"
// @Success      200   {object}  response.Base{data=domain.Wisata}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/wisata/{id} [put]
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

// Delete godoc
// @Summary      Hapus wisata
// @Description  Hapus data wisata berdasarkan ID (admin only)
// @Tags         Wisata
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Wisata ID"
// @Success      200  {object}  response.Base
// @Failure      400  {object}  response.Base
// @Failure      401  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /admin/wisata/{id} [delete]
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