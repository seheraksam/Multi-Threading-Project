package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/Multi-Threading-Project/initializers"
	"github.com/seheraksam/Multi-Threading-Project/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTitle(c *gin.Context) {
	// MongoDB bağlantısı kontrol edelim
	if initializers.MongoClient == nil {
		panic("MongoDB client is nil. Please check MongoDB connection.")
	}

	// Kullanıcıdan gelen URL'yi al
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
		return
	}

	title, err := utils.FetchTitle(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch title"})
		return
	}

	var existingDoc bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = initializers.MongoClient.Database("scrape").Collection("title").FindOne(ctx, bson.M{"url": url}).Decode(&existingDoc)
	if err == nil {
		fmt.Println("🟢 Zaten MongoDB'de var:", url)
		c.JSON(http.StatusOK, gin.H{"message": "Already exists in DB", "url": url, "title": title})
		return
	}

	// **3️⃣ Yeni veriyi MongoDB'ye kaydet**
	_, err = initializers.MongoClient.Database("scrape").Collection("title").InsertOne(ctx, bson.M{
		"url":       url,
		"title":     title,
		"timestamp": time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save in DB"})
		return
	}

	fmt.Println("✅ Yeni veri MongoDB'ye eklendi:", url)
	c.JSON(http.StatusOK, gin.H{"url": url, "title": title})
}
