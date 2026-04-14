package domain

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID          uuid.UUID `db:"id"`
	AdminID     uuid.UUID `db:"admin_id"`
	Name        string    `db:"name"`
	City        string    `db:"city"`
	Category    string    `db:"category"`
	Address     string    `db:"address"`
	Price       float64   `db:"price"`
	Latitude    float64   `db:"latitude"`
	Longitude   float64   `db:"longitude"`
	Description string    `db:"description"`
	Website     string    `db:"website"`
	IsHalal     bool      `db:"is_halal"`
	CreatedAt   time.Time `db:"created_at"`
}