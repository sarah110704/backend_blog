package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database
var DBName = "blog"

// Ambil connection string dari .env atau fallback ke localhost
var MongoString = func() string {
	// Load .env (jika ada)
	_ = godotenv.Load()

	if uri := os.Getenv("MONGOSTRING"); uri != "" {
		return uri
	}
	return "mongodb://localhost:27017" // fallback
}()

func MongoConnect(dbname string) (db *mongo.Database) {
	clientOpts := options.Client().ApplyURI(MongoString)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		fmt.Println("MongoConnect: failed to connect:", err)
		return nil
	}

	// Ping untuk pastikan koneksi sukses
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("MongoConnect: ping failed:", err)
		return nil
	}

	fmt.Println("MongoConnect: connected to MongoDB")
	return client.Database(dbname)
}
