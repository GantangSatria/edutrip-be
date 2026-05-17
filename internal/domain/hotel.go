package domain

import "time"

type Hotel struct {
    ID          int       `db:"id" json:"id"`
    Kota        string    `db:"kota" json:"kota"`
    NamaHotel   string    `db:"nama_hotel" json:"nama_hotel"`
    TipeHotel   string    `db:"tipe_hotel" json:"tipe_hotel"`
    HargaHotel  float64   `db:"harga_hotel" json:"harga_hotel"`
    AlamatHotel string    `db:"alamat_hotel" json:"alamat_hotel"`
    Latitude    float64   `db:"latitude" json:"latitude"`
    Longitude   float64   `db:"longitude" json:"longitude"`
    KetHotel    string    `db:"ket_hotel" json:"ket_hotel"`
    Foto        string    `db:"foto" json:"foto"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
}