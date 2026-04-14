package response

import "github.com/gofiber/fiber/v3"

type Base struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func OK(c fiber.Ctx, message string, data any) error {
	return c.Status(fiber.StatusOK).JSON(Base{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(c fiber.Ctx, message string, data any) error {
	return c.Status(fiber.StatusCreated).JSON(Base{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func BadRequest(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Base{
		Success: false,
		Message: message,
	})
}

func Unauthorized(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Base{
		Success: false,
		Message: message,
	})
}

func NotFound(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(Base{
		Success: false,
		Message: message,
	})
}

func InternalError(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Base{
		Success: false,
		Message: message,
	})
}