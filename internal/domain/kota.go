package domain

import "time"

type Kota struct {
	ID                 int       `db:"id" json:"id"`
	Name               string    `db:"name" json:"name"`
	Image              string    `db:"image" json:"image"`
	Description        string    `db:"description" json:"description"`
	HalalSpotsValue    string    `db:"halal_spots_value" json:"halal_spots_value"`
	HalalSpotsLabel    string    `db:"halal_spots_label" json:"halal_spots_label"`
	MainMosqueTitle    string    `db:"main_mosque_title" json:"main_mosque_title"`
	MainMosqueSubtitle string    `db:"main_mosque_subtitle" json:"main_mosque_subtitle"`
	Features           string    `db:"features" json:"features"` // Use string or jsonb mapping
	TerrainLead        string    `db:"terrain_lead" json:"terrain_lead"`
	TerrainRest        string    `db:"terrain_rest" json:"terrain_rest"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
}
