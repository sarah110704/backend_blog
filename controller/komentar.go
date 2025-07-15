package controller

import (
	"Backend/config"
	"Backend/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllKomentars(ctx context.Context) ([]model.Komentar, error) {
	collection := config.DB.Collection("komentar")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("GetAllKomentars: %v\n", err)
		return nil, err
	}

	var komentars []model.Komentar
	if err := cursor.All(ctx, &komentars); err != nil {
		fmt.Printf("Cursor.All: %v\n", err)
		return nil, err
	}

	return komentars, nil
}

func CreateKomentar(k *model.Komentar) error {
	collection := config.DB.Collection("komentar")

	k.ID = primitive.NewObjectID().Hex()

	_, err := collection.InsertOne(context.TODO(), k)
	if err != nil {
		fmt.Printf("CreateKomentar: %v\n", err)
		return err
	}
	return nil
}

func GetKomentarByID(ctx context.Context, id string) (model.Komentar, error) {
	collection := config.DB.Collection("komentar")
	var komentar model.Komentar

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&komentar)
	if err != nil {
		return komentar, fmt.Errorf("Komentar tidak ditemukan: %v", err)
	}

	return komentar, nil
}

func UpdateKomentarByID(ctx context.Context, id string, data model.Komentar) error {
	collection := config.DB.Collection("komentar")

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"id_artikel": data.IDArtikel,
			"nama":       data.Nama,
			"isi":        data.Isi,
			"tanggal":    data.Tanggal,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("Gagal update komentar: %v", err)
	}

	return nil
}

func DeleteKomentarByID(ctx context.Context, id string) error {
	collection := config.DB.Collection("komentar")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("Gagal menghapus komentar: %v", err)
	}

	return nil
}
