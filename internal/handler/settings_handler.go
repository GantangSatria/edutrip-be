package handler

import (
	"github.com/GantangSatria/edutrip-be/internal/domain"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/GantangSatria/edutrip-be/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type SettingHandler struct {
	svc service.SettingService
}

func NewSettingHandler(svc service.SettingService) *SettingHandler {
	return &SettingHandler{svc}
}

// GetAll godoc
// @Summary      List settings
// @Description  Ambil semua settings
// @Tags         Settings
// @Produce      json
// @Success      200  {object}  response.Base{data=[]domain.Setting}
// @Failure      500  {object}  response.Base
// @Router       /settings [get]
func (h *SettingHandler) GetAll(c fiber.Ctx) error {
	data, err := h.svc.GetAllSettings(c.Context())
	if err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "success", data)
}

// Update godoc
// @Summary      Update setting
// @Description  Perbarui data setting (admin only)
// @Tags         Settings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        key   path      string          true  "Setting Key"
// @Param        body  body      domain.Setting  true  "Data setting"
// @Success      200   {object}  response.Base{data=domain.Setting}
// @Failure      400   {object}  response.Base
// @Failure      401   {object}  response.Base
// @Failure      404   {object}  response.Base
// @Router       /admin/settings/{key} [put]
func (h *SettingHandler) Update(c fiber.Ctx) error {
	key := c.Params("key")
	if key == "" {
		return response.BadRequest(c, "key is required")
	}

	var body domain.Setting
	if err := c.Bind().JSON(&body); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	body.Key = key

	if err := h.svc.UpdateSetting(c.Context(), &body); err != nil {
		return utils.HandleServiceError(c, err)
	}
	return response.OK(c, "setting berhasil diperbarui", body)
}
