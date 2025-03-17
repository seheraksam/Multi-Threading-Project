package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var ProductCollection *mongo.Collection

func ConnecttoMongo() error {
	// MongoDB client ayarı
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB bağlantısı başarısız: %v", err)
	}

	// Bağlantıyı doğrula
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB'ye ping atılamadı: %v", err)
	}

	log.Println("MongoDB'ye başarıyla bağlanıldı!")

	// Koleksiyonun atanması
	ProductCollection = Client.Database("product").Collection("products")
	return nil
}
