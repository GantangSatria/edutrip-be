package domain

import "time"

type RestoranHalal struct {
    ID          int       `db:"id"`
    NamaResto   string    `db:"nama_resto"`
    Kota        string    `db:"kota"`
    AlamatResto string    `db:"alamat_resto"`
    Latitude    float64   `db:"latitude"`
    Longitude   float64   `db:"longitude"`
    KetResto    string    `db:"ket_resto"`
    HargaResto  string    `db:"harga_resto"`
    Foto        string    `db:"foto"`
    CreatedAt   time.Time `db:"created_at"`
}