package bootstrap

import (
	"github.com/GantangSatria/edutrip-be/config"
	"github.com/GantangSatria/edutrip-be/internal/handler"
	"github.com/GantangSatria/edutrip-be/internal/middleware"
	"github.com/GantangSatria/edutrip-be/internal/repository"
	"github.com/GantangSatria/edutrip-be/internal/routes"
	"github.com/GantangSatria/edutrip-be/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func NewApp(cfg *config.Config, db *pgxpool.Pool) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "EduTrip API",
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// ── Repositories ──────────────────────────────────────────────────────────
	adminRepo          := repository.NewAdminRepository(db)
	fasilitasIbadahRepo := repository.NewFasilitasIbadahRepository(db)
	hotelRepo          := repository.NewHotelRepository(db)
	restoranHalalRepo  := repository.NewRestoranHalalRepository(db)
	tokoOlehOlehRepo   := repository.NewTokoOlehOlehRepository(db)
	transportasiRepo   := repository.NewTransportasiRepository(db)
	wisataRepo         := repository.NewWisataRepository(db)

	// ── Services ──────────────────────────────────────────────────────────────
	authService          := service.NewAuthService(adminRepo, cfg.JWTSecret)
	fasilitasIbadahSvc  := service.NewFasilitasIbadahService(fasilitasIbadahRepo)
	hotelSvc            := service.NewHotelService(hotelRepo)
	restoranHalalSvc    := service.NewRestoranHalalService(restoranHalalRepo)
	tokoOlehOlehSvc     := service.NewTokoOlehOlehService(tokoOlehOlehRepo)
	transportasiSvc     := service.NewTransportasiService(transportasiRepo)
	wisataSvc           := service.NewWisataService(wisataRepo)

	// ── Handlers ──────────────────────────────────────────────────────────────
	authHandler          := handler.NewAuthHandler(authService)
	fasilitasIbadahHdlr := handler.NewFasilitasIbadahHandler(fasilitasIbadahSvc)
	hotelHdlr           := handler.NewHotelHandler(hotelSvc)
	restoranHalalHdlr   := handler.NewRestoranHalalHandler(restoranHalalSvc)
	tokoOlehOlehHdlr    := handler.NewTokoOlehOlehHandler(tokoOlehOlehSvc)
	transportasiHdlr    := handler.NewTransportasiHandler(transportasiSvc)
	wisataHdlr          := handler.NewWisataHandler(wisataSvc)
	planDataHdlr        := handler.NewPlanDataHandler(
		wisataSvc,
		hotelSvc,
		restoranHalalSvc,
		fasilitasIbadahSvc,
		tokoOlehOlehSvc,
	)

	// ── Middleware ────────────────────────────────────────────────────────────
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	// ── Routes ────────────────────────────────────────────────────────────────
	routes.Register(
		app,
		authHandler,
		fasilitasIbadahHdlr,
		hotelHdlr,
		restoranHalalHdlr,
		tokoOlehOlehHdlr,
		transportasiHdlr,
		wisataHdlr,
		planDataHdlr,
		authMiddleware,
	)

	return app
}