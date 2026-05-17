package domain

import "time"

type Wisata struct {
    ID            int       `db:"id"`
    Kota          string    `db:"kota"`
    KategoriWisata string   `db:"kategori_wisata"`
    NamaWisata    string    `db:"nama_wisata"`
    TiketWisata   float64   `db:"tiket_wisata"`
    AlamatWisata  string    `db:"alamat_wisata"`
    Latitude      float64   `db:"latitude"`
    Longitude     float64   `db:"longitude"`
    KetWisata     string    `db:"ket_wisata"`
    Foto          string    `db:"foto"`
    CreatedAt     time.Time `db:"created_at"`
}