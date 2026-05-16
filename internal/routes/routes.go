package routes

import (
	"github.com/GantangSatria/edutrip-be/internal/handler"
	"github.com/GantangSatria/edutrip-be/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func Register(
	app *fiber.App,
	authHandler *handler.AuthHandler,

	fasilitasIbadahHandler *handler.FasilitasIbadahHandler,
	hotelHandler           *handler.HotelHandler,
	restoranHalalHandler   *handler.RestoranHalalHandler,
	tokoOlehOlehHandler    *handler.TokoOlehOlehHandler,
	transportasiHandler    *handler.TransportasiHandler,
	wisataHandler          *handler.WisataHandler,

	authMiddleware *middleware.AuthMiddleware,
) {
	api := app.Group("/api/v1")

	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "edutrip api"})
	})

	// ── Public ────────────────────────────────────────────────────────────────
	api.Post("/auth/login", authHandler.Login)

	api.Get("/fasilitas-ibadah", fasilitasIbadahHandler.GetAll)
	api.Get("/fasilitas-ibadah/:id", fasilitasIbadahHandler.GetByID)

	api.Get("/hotel", hotelHandler.GetAll)
	api.Get("/hotel/:id", hotelHandler.GetByID)

	api.Get("/restoran-halal", restoranHalalHandler.GetAll)
	api.Get("/restoran-halal/:id", restoranHalalHandler.GetByID)

	api.Get("/toko-oleh-oleh", tokoOlehOlehHandler.GetAll)
	api.Get("/toko-oleh-oleh/:id", tokoOlehOlehHandler.GetByID)

	api.Get("/transportasi", transportasiHandler.GetAll)
	api.Get("/transportasi/:id", transportasiHandler.GetByID)

	api.Get("/wisata", wisataHandler.GetAll)
	api.Get("/wisata/:id", wisataHandler.GetByID)

	// ── CMS — protected, hanya admin ─────────────────────────────────────────
	admin := api.Group("/admin", authMiddleware.Protected())

	admin.Post("/fasilitas-ibadah", fasilitasIbadahHandler.Create)
	admin.Put("/fasilitas-ibadah/:id", fasilitasIbadahHandler.Update)
	admin.Delete("/fasilitas-ibadah/:id", fasilitasIbadahHandler.Delete)

	admin.Post("/hotel", hotelHandler.Create)
	admin.Put("/hotel/:id", hotelHandler.Update)
	admin.Delete("/hotel/:id", hotelHandler.Delete)

	admin.Post("/restoran-halal", restoranHalalHandler.Create)
	admin.Put("/restoran-halal/:id", restoranHalalHandler.Update)
	admin.Delete("/restoran-halal/:id", restoranHalalHandler.Delete)

	admin.Post("/toko-oleh-oleh", tokoOlehOlehHandler.Create)
	admin.Put("/toko-oleh-oleh/:id", tokoOlehOlehHandler.Update)
	admin.Delete("/toko-oleh-oleh/:id", tokoOlehOlehHandler.Delete)

	admin.Post("/transportasi", transportasiHandler.Create)
	admin.Put("/transportasi/:id", transportasiHandler.Update)
	admin.Delete("/transportasi/:id", transportasiHandler.Delete)

	admin.Post("/wisata", wisataHandler.Create)
	admin.Put("/wisata/:id", wisataHandler.Update)
	admin.Delete("/wisata/:id", wisataHandler.Delete)
}