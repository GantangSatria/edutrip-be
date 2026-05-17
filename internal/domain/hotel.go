package domain

import "time"

type Hotel struct {
    ID          int       `db:"id"`
    Kota        string    `db:"kota"`
    NamaHotel   string    `db:"nama_hotel"`
    TipeHotel   string    `db:"tipe_hotel"`
    HargaHotel  float64   `db:"harga_hotel"`
    AlamatHotel string    `db:"alamat_hotel"`
    Latitude    float64   `db:"latitude"`
    Longitude   float64   `db:"longitude"`
    KetHotel    string    `db:"ket_hotel"`
    Foto        string    `db:"foto"`
    CreatedAt   time.Time `db:"created_at"`
}