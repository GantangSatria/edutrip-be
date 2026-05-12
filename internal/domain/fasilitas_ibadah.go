package domain

import "time"

type FasilitasIbadah struct {
    ID             int       `db:"id"`
    Kota           string    `db:"kota"`
    TipeFas        string    `db:"tipe_fas"`
    NamaFasIbadah  string    `db:"nama_fas_ibadah"`
    LokasiFasIbadah string   `db:"lokasi_fas_ibadah"`
    Latitude       float64   `db:"latitude"`
    Longitude      float64   `db:"longitude"`
    Foto           string    `db:"foto"`
    CreatedAt      time.Time `db:"created_at"`
}