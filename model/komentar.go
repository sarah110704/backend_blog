package model

import "time"

type Komentar struct {
    ID        string    `json:"_id" bson:"_id"` // contoh: "c001"
    IDArtikel string    `json:"id_artikel" bson:"id_artikel"`
    Nama      string    `json:"nama" bson:"nama"`
    Isi       string    `json:"isi" bson:"isi"`
    Tanggal   time.Time `json:"tanggal" bson:"tanggal"`
}
