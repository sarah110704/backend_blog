package model

type Kategori struct {
    ID        string `json:"_id" bson:"_id"` // contoh: "k001"
    Nama      string `json:"nama" bson:"nama"`
    Deskripsi string `json:"deskripsi" bson:"deskripsi"`
}
