package model

import "time"

type Artikel struct {
	ID         string    `json:"id" bson:"_id"`
	Judul      string    `json:"judul" bson:"judul"`
	Isi        string    `json:"isi" bson:"isi"`
	Tanggal    time.Time `json:"tanggal" bson:"tanggal"`
	IDPenulis  string    `json:"id_penulis" bson:"id_penulis"`
	IDKategori string    `json:"id_kategori" bson:"id_kategori"`
}
