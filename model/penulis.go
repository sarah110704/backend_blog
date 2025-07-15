package model

type Penulis struct {
    ID    string `json:"_id" bson:"_id"` // contoh: "p001"
    Nama  string `json:"nama" bson:"nama"`
    Email string `json:"email" bson:"email"`
    Bio   string `json:"bio" bson:"bio"`
}
