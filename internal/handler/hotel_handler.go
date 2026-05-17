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

// GetAll godoc
// @Summary      List hotel
// @Description  Ambil semua hotel. Filter opsional dengan query param kota dan/atau tipe.
// @Tags         Hotel
// @Produce      json
// @Param        kota  query     string  false  "Filter by kota"
// @Param        tipe  query     string  false  "Filter by tipe hotel"
// @Success      200   {object}  response.Base{data=[]domain.Hotel}
// @Failure      500   {object}  response.Base
// @Router       /hotel [get]
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

// GetByID godoc
// @Summary      Detail hotel
// @Description  Ambil satu hotel berdasarkan ID
// @Tags         Hotel
// @Produce      json
// @Param        id   path      int  true  "Hotel ID"
// @Success      200  {object}  response.Base{data=domain.Hotel}
// @Failure      400  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /hotel/{id} [get]
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

// Create godoc
// @Summary      Tambah hotel
// @Description  Tambah data hotel baru (admin only)
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      domain.Hotel  true  "Data hotel"
// @Success      201   {object}  response.Base{data=domain.Hotel}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Router       /admin/hotel [post]
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

// Update godoc
// @Summary      Update hotel
// @Description  Perbarui data hotel berdasarkan ID (admin only)
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int           true  "Hotel ID"
// @Param        body  body      domain.Hotel  true  "Data hotel"
// @Success      200   {object}  response.Base{data=domain.Hotel}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/hotel/{id} [put]
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

// Delete godoc
// @Summary      Hapus hotel
// @Description  Hapus data hotel berdasarkan ID (admin only)
// @Tags         Hotel
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Hotel ID"
// @Success      200  {object}  response.Base
// @Failure      400  {object}  response.Base
// @Failure      401  {object}  response.Base
// @Failure      404  {object}  response.Base
// @Router       /admin/hotel/{id} [delete]
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