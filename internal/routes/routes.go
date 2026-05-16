package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/GantangSatria/edutrip-be/internal/handler"
	"github.com/GantangSatria/edutrip-be/internal/middleware"
)

func Register(
	app *fiber.App,
	authHandler *handler.AuthHandler,

	authMiddleware *middleware.AuthMiddleware,
) {
	api := app.Group("/api/v1")

	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "edutrip api"})
	})

	// Public
	api.Post("/auth/login", authHandler.Login)
	

	// CMS — protected, hanya admin
	admin := api.Group("/admin", authMiddleware.Protected())

}