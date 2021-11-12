package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/buker/go-api-starter/internal/models"
	db "github.com/buker/go-api-starter/internal/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var timeEntriesCollection *mongo.Collection = db.OpenCollection(db.Client, "timeEntries")

//Insert time entry into timeEntries collection
func InsertTimeEntry() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var timeEntry models.TimeEntry
		var foundUser models.User

		if err := c.BindJSON(&timeEntry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"user_id": timeEntry.User_id}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		timeEntry.ID = primitive.NewObjectID()
		timeEntry.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		timeEntry.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		timeEntry.Time_start, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		timeEntry.Time_end, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		timeEntry.User_id = foundUser.ID
		resultInsertionNumber, insertErr := timeEntriesCollection.InsertOne(ctx, timeEntry)
		if insertErr != nil {
			msg := fmt.Sprintf("Time entry item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}
