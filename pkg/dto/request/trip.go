package request

type CreateTrip struct {
	Title            string  `json:"title"             validate:"required"`
	Description      string  `json:"description"`
	City             string  `json:"city"              validate:"required"`
	Theme            string  `json:"theme"`
	DurationDays     int     `json:"duration_days"     validate:"required,min=1"`
	Price            float64 `json:"price"             validate:"required,min=0"`
	Facilities       string  `json:"facilities"`
	ItinerarySummary string  `json:"itinerary_summary"`
	Status           string  `json:"status"            validate:"omitempty,oneof=draft published archived"`
}

type UpdateTrip struct {
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	City             string  `json:"city"`
	Theme            string  `json:"theme"`
	DurationDays     int     `json:"duration_days"`
	Price            float64 `json:"price"`
	Facilities       string  `json:"facilities"`
	ItinerarySummary string  `json:"itinerary_summary"`
	Status           string  `json:"status"            validate:"omitempty,oneof=draft published archived"`
}