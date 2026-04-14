package bootstrap

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GantangSatria/edutrip-be/config"
)

func NewDB(cfg *config.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), cfg.DBUrl)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	log.Println("database connected")
	return pool
}