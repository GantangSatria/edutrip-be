package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/GantangSatria/edutrip-be/pkg/dto/request"
	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
)

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req request.Login
	if err := c.Bind().JSON(&req); err != nil {
		return response.BadRequest(c, "invalid request body")
	}

	token, err := h.svc.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	return response.OK(c, "login successful", fiber.Map{"token": token})
}