package utils

import (
	"errors"
	"strconv"

	"github.com/GantangSatria/edutrip-be/pkg/dto/response"
	"github.com/gofiber/fiber/v3"
)

// parseID mengambil path param "id" dan mengkonversinya ke int.
func parseID(c fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id")
	}
	return id, nil
}

// handleServiceError memetakan error repository ke HTTP response yang sesuai.
func handleServiceError(c fiber.Ctx, err error) error {
	if errors.Is(err, ErrNotFound) {
		return response.NotFound(c, "data tidak ditemukan")
	}
	return response.InternalError(c, "internal server error")
}