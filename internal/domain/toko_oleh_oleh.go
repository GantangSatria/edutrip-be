package domain

import "time"

type TokoOlehOleh struct {
    ID            int       `db:"id"`
    Kota          string    `db:"kota"`
    NamaBelanja   string    `db:"nama_belanja"`
    JenisBelanja  string    `db:"jenis_belanja"`
    KetBelanja    string    `db:"ket_belanja"`
    AlamatBelanja string    `db:"alamat_belanja"`
    Latitude      float64   `db:"latitude"`
    Longitude     float64   `db:"longitude"`
    Foto          string    `db:"foto"`
    CreatedAt     time.Time `db:"created_at"`
}