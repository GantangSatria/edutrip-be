package domain

import "time"

type Transportasi struct {
    ID                   int       `db:"id" json:"id"`
    JenisTransportasi    string    `db:"jenis_transportasi" json:"jenis_transportasi"`
    NamaTransportasi     string    `db:"nama_transportasi" json:"nama_transportasi"`
    Rute                 string    `db:"rute" json:"rute"`
    KodeBandara          string    `db:"kode_bandara" json:"kode_bandara"`
    HargaTransportasiIDR float64   `db:"harga_transportasi_idr" json:"harga_transportasi_idr"`
    KetTransportasi      string    `db:"ket_transportasi" json:"ket_transportasi"`
    Foto                 string    `db:"foto" json:"foto"`
    CreatedAt            time.Time `db:"created_at" json:"created_at"`
}