package request

type CreateLocation struct {
	Name        string  `json:"name"        validate:"required"`
	City        string  `json:"city"        validate:"required"`
	Category    string  `json:"category"    validate:"required"`
	Address     string  `json:"address"`
	Price       float64 `json:"price"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
	Website     string  `json:"website"`
	IsHalal     bool    `json:"is_halal"`
}

type UpdateLocation struct {
	Name        string  `json:"name"`
	City        string  `json:"city"`
	Category    string  `json:"category"`
	Address     string  `json:"address"`
	Price       float64 `json:"price"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
	Website     string  `json:"website"`
	IsHalal     bool    `json:"is_halal"`
}