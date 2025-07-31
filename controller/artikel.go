package controller

import (
	"Backend/config"
	"Backend/model"
	"context"
	"fmt"

	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ✅ Ambil semua artikel dengan Nama Kategori dan Penulis
func GetAllArtikels(ctx context.Context) ([]model.Artikel, error) {
	collection := config.DB.Collection("artikel")

	// Ambil data artikel
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("GetAllArtikels: %v\n", err)
		return nil, err
	}

	var artikels []model.Artikel
	if err := cursor.All(ctx, &artikels); err != nil {
		fmt.Printf("Cursor.All: %v\n", err)
		return nil, err
	}

	// Menambahkan Nama Kategori dan Penulis ke setiap artikel
	for i, artikel := range artikels {
		// Ambil nama kategori dari collection kategori
		kategoriCollection := config.DB.Collection("kategori")
		var kategori model.Kategori
		err := kategoriCollection.FindOne(ctx, bson.M{"_id": artikel.IDKategori}).Decode(&kategori)
		if err == nil {
			artikels[i].KategoriNama = kategori.Nama
		}

		// Ambil nama penulis dari collection penulis
		penulisCollection := config.DB.Collection("penulis")
		var penulis model.Penulis
		err = penulisCollection.FindOne(ctx, bson.M{"_id": artikel.IDPenulis}).Decode(&penulis)
		if err == nil {
			artikels[i].PenulisNama = penulis.Nama
		}
	}

	return artikels, nil
}

// ✅ Tambah artikel baru
func CreateArtikel(artikel *model.Artikel) error {
	collection := config.DB.Collection("artikel")

	// Buat ID otomatis
	artikel.ID = primitive.NewObjectID().Hex()

	_, err := collection.InsertOne(context.TODO(), artikel)
	if err != nil {
		fmt.Println("CreateArtikel:", err)
		return err
	}

	return nil
}

// ✅ Ambil artikel berdasarkan ID
func GetArtikelByID(ctx context.Context, id string) (model.Artikel, error) {
	var artikel model.Artikel
	collection := config.DB.Collection("artikel")

	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&artikel)
	if err != nil {
		return artikel, fmt.Errorf("data tidak ditemukan: %v", err)
	}

	return artikel, nil
}

// ✅ Update artikel berdasarkan ID
func UpdateArtikelByID(ctx context.Context, id string, data model.Artikel) error {
	collection := config.DB.Collection("artikel")

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"judul":       data.Judul,
			"isi":         data.Isi,
			"tanggal":     data.Tanggal,
			"id_penulis":  data.IDPenulis,
			"id_kategori": data.IDKategori,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// ✅ Hapus artikel berdasarkan ID
func DeleteArtikelByID(ctx context.Context, id string) error {
	collection := config.DB.Collection("artikel")

	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// ✅ Mengekspor artikel ke Excel
func ExportArtikelsToExcel(ctx context.Context) ([]byte, error) {
	artikels, err := GetAllArtikels(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data artikel: %v", err)
	}

	excel := excelize.NewFile()
	sheet := "Sheet1"
	excel.SetSheetName("Sheet1", sheet)

	// Set Header
	excel.SetCellValue(sheet, "A1", "Judul")
	excel.SetCellValue(sheet, "B1", "Isi")
	excel.SetCellValue(sheet, "C1", "Kategori")
	excel.SetCellValue(sheet, "D1", "Penulis")

	// Isi data artikel
	for i, artikel := range artikels {
		row := i + 2
		excel.SetCellValue(sheet, fmt.Sprintf("A%d", row), artikel.Judul)
		excel.SetCellValue(sheet, fmt.Sprintf("B%d", row), artikel.Isi)
		excel.SetCellValue(sheet, fmt.Sprintf("C%d", row), artikel.KategoriNama) // Nama kategori
		excel.SetCellValue(sheet, fmt.Sprintf("D%d", row), artikel.PenulisNama)  // Nama penulis
	}

	// Menyimpan ke buffer
	buffer, err := excel.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("gagal menulis file Excel: %v", err)
	}

	return buffer.Bytes(), nil
}
