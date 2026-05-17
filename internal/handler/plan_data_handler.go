package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type PlanDataHandler struct {
	wisataSvc          service.WisataService
	hotelSvc           service.HotelService
	restoranHalalSvc   service.RestoranHalalService
	fasilitasIbadahSvc service.FasilitasIbadahService
	tokoOlehOlehSvc    service.TokoOlehOlehService
}

func NewPlanDataHandler(
	wisataSvc service.WisataService,
	hotelSvc service.HotelService,
	restoranHalalSvc service.RestoranHalalService,
	fasilitasIbadahSvc service.FasilitasIbadahService,
	tokoOlehOlehSvc service.TokoOlehOlehService,
) *PlanDataHandler {
	return &PlanDataHandler{
		wisataSvc:          wisataSvc,
		hotelSvc:           hotelSvc,
		restoranHalalSvc:   restoranHalalSvc,
		fasilitasIbadahSvc: fasilitasIbadahSvc,
		tokoOlehOlehSvc:    tokoOlehOlehSvc,
	}
}

// GetPlanData godoc
// @Summary      Aggregator data perencanaan trip
// @Description  Ambil semua data yang dibutuhkan untuk merencanakan trip dalam satu request. Filter opsional dengan query param kota.
// @Tags         Plan Data
// @Produce      json
// @Param        kota  query     string  false  "Filter semua data by kota"
// @Success      200   {object}  response.Base{data=response.PlanData}
// @Failure      500   {object}  response.Base
// @Router       /plan-data [get]
func (h *PlanDataHandler) GetPlanData(c fiber.Ctx) error {
	ctx := c.Context()
	kota := c.Query("kota")

	type result struct {
		key string
		val any
		err error
	}

	ch := make(chan result, 5)

	go func() {
		if kota != "" {
			data, err := h.wisataSvc.GetByKota(ctx, kota)
			ch <- result{"wisata", data, err}
		} else {
			data, err := h.wisataSvc.GetAll(ctx)
			ch <- result{"wisata", data, err}
		}
	}()

	go func() {
		if kota != "" {
			data, err := h.hotelSvc.GetByKota(ctx, kota)
			ch <- result{"hotel", data, err}
		} else {
			data, err := h.hotelSvc.GetAll(ctx)
			ch <- result{"hotel", data, err}
		}
	}()

	go func() {
		if kota != "" {
			data, err := h.restoranHalalSvc.GetByKota(ctx, kota)
			ch <- result{"restoran", data, err}
		} else {
			data, err := h.restoranHalalSvc.GetAll(ctx)
			ch <- result{"restoran", data, err}
		}
	}()

	go func() {
		if kota != "" {
			data, err := h.fasilitasIbadahSvc.GetByKota(ctx, kota)
			ch <- result{"fasilitas_ibadah", data, err}
		} else {
			data, err := h.fasilitasIbadahSvc.GetAll(ctx)
			ch <- result{"fasilitas_ibadah", data, err}
		}
	}()

	go func() {
		if kota != "" {
			data, err := h.tokoOlehOlehSvc.GetByKota(ctx, kota)
			ch <- result{"toko_oleh_oleh", data, err}
		} else {
			data, err := h.tokoOlehOlehSvc.GetAll(ctx)
			ch <- result{"toko_oleh_oleh", data, err}
		}
	}()

	planData := fiber.Map{
		"wisata":           []any{},
		"hotel":            []any{},
		"restoran":         []any{},
		"fasilitas_ibadah": []any{},
		"toko_oleh_oleh":   []any{},
	}

	for i := 0; i < 5; i++ {
		r := <-ch
		if r.err != nil {
			return utils.HandleServiceError(c, r.err)
		}
		planData[r.key] = r.val
	}

	return response.OK(c, "success", planData)
}