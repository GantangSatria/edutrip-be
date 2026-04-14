package main

import (
	"log"

	"github.com/GantangSatria/edutrip-be/internal/bootstrap"
	"github.com/GantangSatria/edutrip-be/config"
)

func main() {
	cfg := config.Load()

	db := bootstrap.NewDB(cfg)
	defer db.Close()

	app := bootstrap.NewApp(cfg, db)

	log.Printf("Server running on port %s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}