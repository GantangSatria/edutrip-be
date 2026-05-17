package domain

import "time"

type FasilitasIbadah struct {
    ID              int       `db:"id" json:"id"`
    Kota            string    `db:"kota" json:"kota"`
    TipeFas         string    `db:"tipe_fas" json:"tipe_fas"`
    NamaFasIbadah   string    `db:"nama_fas_ibadah" json:"nama_fas_ibadah"`
    LokasiFasIbadah string    `db:"lokasi_fas_ibadah" json:"lokasi_fas_ibadah"`
    Latitude        float64   `db:"latitude" json:"latitude"`
    Longitude       float64   `db:"longitude" json:"longitude"`
    Foto            string    `db:"foto" json:"foto"`
    CreatedAt       time.Time `db:"created_at" json:"created_at"`
}