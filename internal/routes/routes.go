package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/GantangSatria/edutrip-be/internal/handler"
	"github.com/GantangSatria/edutrip-be/internal/middleware"
)

func Register(
	app *fiber.App,
	authHandler *handler.AuthHandler,
	tripHandler *handler.TripHandler,
	locationHandler *handler.LocationHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	api := app.Group("/api/v1")

	// Public
	api.Post("/auth/login", authHandler.Login)
	
	api.Get("/trips", tripHandler.GetAll)
	api.Get("/trips/:id", tripHandler.GetByID)
	api.Get("/locations", locationHandler.GetAll)
	api.Get("/locations/:id", locationHandler.GetByID)

	// CMS — protected, hanya admin
	admin := api.Group("/admin", authMiddleware.Protected())

	admin.Post("/trips", tripHandler.Create)
	admin.Put("/trips/:id", tripHandler.Update)
	admin.Delete("/trips/:id", tripHandler.Delete)

	admin.Post("/locations", locationHandler.Create)
	admin.Put("/locations/:id", locationHandler.Update)
	admin.Delete("/locations/:id", locationHandler.Delete)
}