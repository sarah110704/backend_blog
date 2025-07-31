package controller

import (
	"Backend/config"
	"Backend/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// GetAllKategoris - Ambil semua kategori
func GetAllKategoris(ctx context.Context) ([]model.Kategori, error) {
	collection := config.DB.Collection("kategori")
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetAllKategoris: %v\n", err)
		return nil, err
	}

	var data []model.Kategori
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Printf("GetAllKategoris: %v\n", err)
		return nil, err
	}

	return data, nil
}

// CreateKategori - Menambahkan kategori baru
func CreateKategori(kategori *model.Kategori) error {
	collection := config.DB.Collection("kategori")

	_, err := collection.InsertOne(context.TODO(), kategori)
	if err != nil {
		fmt.Println("CreateKategori:", err)
		return err
	}

	return nil
}

// GetKategoriByID - Mengambil kategori berdasarkan ID
func GetKategoriByID(ctx context.Context, id string) (model.Kategori, error) {
	var kategori model.Kategori
	collection := config.DB.Collection("kategori")

	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&kategori)
	if err != nil {
		return kategori, fmt.Errorf("data tidak ditemukan: %v", err)
	}

	return kategori, nil
}

// UpdateKategoriByID - Update kategori berdasarkan ID
func UpdateKategoriByID(ctx context.Context, id string, data model.Kategori) error {
	collection := config.DB.Collection("kategori")

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":      data.Nama,
			"deskripsi": data.Deskripsi,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteKategoriByID - Hapus kategori berdasarkan ID
func DeleteKategoriByID(ctx context.Context, id string) error {
	collection := config.DB.Collection("kategori")

	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
