// @title           EduTrip API
// @version         1.0
// @description     API untuk aplikasi EduTrip — hotel, wisata, transportasi, restoran halal, toko oleh-oleh, dan fasilitas ibadah.
// @termsOfService  http://swagger.io/terms/
 
// @contact.name   EduTrip Dev
// @contact.email  admin@edutrip.com
 
// @host      localhost:8080
// @BasePath  /api/v1
 
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format: Bearer {token}

package main

import (
	"log"

	"github.com/GantangSatria/edutrip-be/internal/bootstrap"
	"github.com/GantangSatria/edutrip-be/config"
	_ "github.com/GantangSatria/edutrip-be/docs"
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