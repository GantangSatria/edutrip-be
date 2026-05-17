package domain

import "time"

type Transportasi struct {
    ID                   int       `db:"id"`
    JenisTransportasi    string    `db:"jenis_transportasi"`
    NamaTransportasi     string    `db:"nama_transportasi"`
    Rute                 string    `db:"rute"`
    KodeBandara          string    `db:"kode_bandara"`
    HargaTransportasiIDR float64   `db:"harga_transportasi_idr"`
    KetTransportasi      string    `db:"ket_transportasi"`
    Foto                 string    `db:"foto"`
    CreatedAt            time.Time `db:"created_at"`
}