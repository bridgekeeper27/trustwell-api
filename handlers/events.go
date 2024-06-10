package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/bridgekeeper27/trustwell-api/database"
	"github.com/bridgekeeper27/trustwell-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = primitive.NewObjectID()
	event.CreatedAt = time.Now()
	event.CreatedBy = c.MustGet("userID").(string)
	event.IsDeleted = false

	collection := database.GetCollection("events")
	_, err := collection.InsertOne(context.Background(), event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("userID").(string)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := database.GetCollection("events")
	filter := bson.M{"_id": objectID, "createdBy": userID}

	update := bson.M{"$set": bson.M{"isDeleted": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}

func GetEvent(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("userID").(string)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := database.GetCollection("events")
	filter := bson.M{"_id": objectID, "createdBy": userID, "isDeleted": false}

	var event models.Event
	err = collection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		}
		return
	}

	c.JSON(http.StatusOK, event)
}

func ListEvents(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	collection := database.GetCollection("events")
	filter := bson.M{"createdBy": userID, "isDeleted": false}
	opts := options.Find().SetSort(bson.M{"createdAt": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list events"})
		return
	}
	defer cursor.Close(context.Background())

	var events []models.Event
	if err = cursor.All(context.Background(), &events); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode events"})
		return
	}

	c.JSON(http.StatusOK, events)
}
