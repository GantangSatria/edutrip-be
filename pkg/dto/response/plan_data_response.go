package response

import "github.com/GantangSatria/edutrip-be/internal/domain"

type PlanData struct {
	Wisata          []domain.Wisata          `json:"wisata"`
	Hotel           []domain.Hotel           `json:"hotel"`
	Restoran        []domain.RestoranHalal   `json:"restoran"`
	FasilitasIbadah []domain.FasilitasIbadah `json:"fasilitas_ibadah"`
	TokoOlehOleh   []domain.TokoOlehOleh    `json:"toko_oleh_oleh"`
}