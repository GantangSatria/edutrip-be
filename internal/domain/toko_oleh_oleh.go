package domain

import "time"

type TokoOlehOleh struct {
    ID            int       `db:"id" json:"id"`
    Kota          string    `db:"kota" json:"kota"`
    NamaBelanja   string    `db:"nama_belanja" json:"nama_belanja"`
    JenisBelanja  string    `db:"jenis_belanja" json:"jenis_belanja"`
    KetBelanja    string    `db:"ket_belanja" json:"ket_belanja"`
    AlamatBelanja string    `db:"alamat_belanja" json:"alamat_belanja"`
    Latitude      float64   `db:"latitude" json:"latitude"`
    Longitude     float64   `db:"longitude" json:"longitude"`
    Foto          string    `db:"foto" json:"foto"`
    CreatedAt     time.Time `db:"created_at" json:"created_at"`
}