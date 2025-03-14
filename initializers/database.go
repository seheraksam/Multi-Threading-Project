package initializers

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var RedisClient *redis.Client

// MongoDB'ye bağlantı kuran fonksiyon
func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB bağlantısı başarısız: %v", err)
	}

	// MongoDB bağlantısının başarılı olup olmadığını kontrol et
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB'ye ping atılamadı: %v", err)
	}

	log.Println("Connected to MongoDB")
}

// Redis'e bağlantı kuran fonksiyon
func ConnectRedis() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println(rdb)
	rdb.Set(ctx, "foo", "bar", 0)
	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar
	fmt.Println("Connected to Redis!")
}
