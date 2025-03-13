package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/Multi-Threading-Project/initializers"
	"github.com/seheraksam/Multi-Threading-Project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = initializers.MongoClient.Database("scraperDB").Collection("titles")

func GetTitle(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL param is required"})
		return
	}

	title, err := utils.FetchTitle(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch title"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, bson.M{"url": url, "title": title, "timestamp": time.Now()})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save in DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url, "title": title})
}
