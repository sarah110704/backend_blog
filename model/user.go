package model

type User struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Nama     string `json:"nama" bson:"nama"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}
