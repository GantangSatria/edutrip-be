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
	fiberSwagger "github.com/gofiber/contrib/v3/swaggo"

	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func NewApp(cfg *config.Config, db *pgxpool.Pool) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "EduTrip API",
	})

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://edutrip-japan.vercel.app",
		},
		AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	// ── Repositories ──────────────────────────────────────────────────────────
	adminRepo          := repository.NewAdminRepository(db)
	fasilitasIbadahRepo := repository.NewFasilitasIbadahRepository(db)
	hotelRepo          := repository.NewHotelRepository(db)
	restoranHalalRepo  := repository.NewRestoranHalalRepository(db)
	tokoOlehOlehRepo   := repository.NewTokoOlehOlehRepository(db)
	transportasiRepo   := repository.NewTransportasiRepository(db)
	wisataRepo         := repository.NewWisataRepository(db)
	settingRepo        := repository.NewSettingRepository(db)
	kotaRepo           := repository.NewKotaRepository(db)

	// ── Services ──────────────────────────────────────────────────────────────
	authService          := service.NewAuthService(adminRepo, cfg.JWTSecret)
	fasilitasIbadahSvc  := service.NewFasilitasIbadahService(fasilitasIbadahRepo)
	hotelSvc            := service.NewHotelService(hotelRepo)
	restoranHalalSvc    := service.NewRestoranHalalService(restoranHalalRepo)
	tokoOlehOlehSvc     := service.NewTokoOlehOlehService(tokoOlehOlehRepo)
	transportasiSvc     := service.NewTransportasiService(transportasiRepo)
	wisataSvc           := service.NewWisataService(wisataRepo)
	settingSvc          := service.NewSettingService(settingRepo)
	kotaSvc             := service.NewKotaService(kotaRepo)

	// ── Handlers ──────────────────────────────────────────────────────────────
	authHandler          := handler.NewAuthHandler(authService)
	fasilitasIbadahHdlr := handler.NewFasilitasIbadahHandler(fasilitasIbadahSvc)
	hotelHdlr           := handler.NewHotelHandler(hotelSvc)
	restoranHalalHdlr   := handler.NewRestoranHalalHandler(restoranHalalSvc)
	tokoOlehOlehHdlr    := handler.NewTokoOlehOlehHandler(tokoOlehOlehSvc)
	transportasiHdlr    := handler.NewTransportasiHandler(transportasiSvc)
	wisataHdlr          := handler.NewWisataHandler(wisataSvc)
	settingHdlr         := handler.NewSettingHandler(settingSvc)
	kotaHdlr            := handler.NewKotaHandler(kotaSvc)
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
		settingHdlr,
		kotaHdlr,
		planDataHdlr,
		authMiddleware,
	)

	return app
}