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
)

func NewApp(cfg *config.Config, db *pgxpool.Pool) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "EduTrip API",
	})

	// Repositories
	adminRepo    := repository.NewAdminRepository(db)
 
	// Services
	authService     := service.NewAuthService(adminRepo, cfg.JWTSecret)
 
	// Handlers
	authHandler     := handler.NewAuthHandler(authService)
 
	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)
 
	// Routes
	routes.Register(app, authHandler, authMiddleware)
 
	return app
}