package model

type User struct {
	ID       string `json:"id" bson:"_id"`
	Nama     string `json:"nama" bson:"nama"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
