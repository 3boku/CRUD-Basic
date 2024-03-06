package services

import (
	"context"
	"crud-server/entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Post(c *gin.Context, db *mongo.Client) {
	var feed entities.FeedDTO

	err := c.BindJSON(&feed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	collection := db.Database("study").Collection("feed")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.InsertOne(ctx, feed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}