package controller

import (
	"Backend/config"
	"Backend/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllPenulis - Ambil semua data penulis
func GetAllPenulis(ctx context.Context) ([]model.Penulis, error) {
	collection := config.DB.Collection("penulis")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("GetAllPenulis:", err)
		return nil, err
	}

	var data []model.Penulis
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllPenulis decode:", err)
		return nil, err
	}

	return data, nil
}

// CreatePenulis - Menambahkan penulis baru
func CreatePenulis(p *model.Penulis) error {
	collection := config.DB.Collection("penulis")
	p.ID = primitive.NewObjectID().Hex()

	_, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		fmt.Println("CreatePenulis:", err)
		return err
	}

	return nil
}

// GetPenulisByID - Mengambil penulis berdasarkan ID
func GetPenulisByID(ctx context.Context, id string) (model.Penulis, error) {
	var p model.Penulis
	collection := config.DB.Collection("penulis")

	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&p)
	if err != nil {
		return p, fmt.Errorf("penulis tidak ditemukan: %v", err)
	}

	return p, nil
}

// UpdatePenulisByID - Update data penulis berdasarkan ID
func UpdatePenulisByID(ctx context.Context, id string, data model.Penulis) error {
	collection := config.DB.Collection("penulis")

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":  data.Nama,
			"email": data.Email,
			"bio":   data.Bio,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

// DeletePenulisByID - Hapus penulis berdasarkan ID
func DeletePenulisByID(ctx context.Context, id string) error {
	collection := config.DB.Collection("penulis")

	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
