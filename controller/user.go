package controller

import (
	"Backend/config"
	"Backend/model"
	"Backend/utils"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(ctx context.Context, user *model.User) error {
	collection := config.DB.Collection("users")

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("RegisterUser: gagal hash password:", err)
		return fmt.Errorf("gagal hash password: %v", err)
	}
	user.Password = string(hashed)
	user.ID = primitive.NewObjectID().Hex()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("RegisterUser: gagal insert user:", err)
		return fmt.Errorf("gagal insert user: %v", err)
	}

	log.Println("RegisterUser: user berhasil disimpan:", user)
	return nil
}

func LoginUser(ctx context.Context, email string, password string) (string, error) {
	collection := config.DB.Collection("users")

	var user model.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Println("LoginUser: user tidak ditemukan:", err)
		return "", fmt.Errorf("user tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("LoginUser: password salah:", err)
		return "", fmt.Errorf("password salah")
	}

	token, err := utils.GenerateJWT(email)
	if err != nil {
		return "", fmt.Errorf("gagal buat token: %v", err)
	}

	return token, nil
}
