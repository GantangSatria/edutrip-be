package domain

import "time"

type RestoranHalal struct {
    ID          int       `db:"id" json:"id"`
    NamaResto   string    `db:"nama_resto" json:"nama_resto"`
    Kota        string    `db:"kota" json:"kota"`
    AlamatResto string    `db:"alamat_resto" json:"alamat_resto"`
    Latitude    float64   `db:"latitude" json:"latitude"`
    Longitude   float64   `db:"longitude" json:"longitude"`
    KetResto    string    `db:"ket_resto" json:"ket_resto"`
    HargaResto  string    `db:"harga_resto" json:"harga_resto"`
    Foto        string    `db:"foto" json:"foto"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
}