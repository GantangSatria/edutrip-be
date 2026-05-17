package domain

import "time"

type Wisata struct {
    ID             int       `db:"id" json:"id"`
    Kota           string    `db:"kota" json:"kota"`
    KategoriWisata string    `db:"kategori_wisata" json:"kategori_wisata"`
    NamaWisata     string    `db:"nama_wisata" json:"nama_wisata"`
    TiketWisata    float64   `db:"tiket_wisata" json:"tiket_wisata"`
    AlamatWisata   string    `db:"alamat_wisata" json:"alamat_wisata"`
    Latitude       float64   `db:"latitude" json:"latitude"`
    Longitude      float64   `db:"longitude" json:"longitude"`
    KetWisata      string    `db:"ket_wisata" json:"ket_wisata"`
    Foto           string    `db:"foto" json:"foto"`
    CreatedAt      time.Time `db:"created_at" json:"created_at"`
}