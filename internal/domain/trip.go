package domain

import (
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	ID               uuid.UUID `db:"id"`
	AdminID          uuid.UUID `db:"admin_id"`
	Title            string    `db:"title"`
	Description      string    `db:"description"`
	City             string    `db:"city"`
	Theme            string    `db:"theme"`
	DurationDays     int       `db:"duration_days"`
	Price            float64   `db:"price"`
	Facilities       string    `db:"facilities"`
	ItinerarySummary string    `db:"itinerary_summary"`
	Status           string    `db:"status"`
	CreatedAt        time.Time `db:"created_at"`
}