package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/buker/TimeGladiator/internal/models"
	db "github.com/buker/TimeGladiator/internal/repository/mongodb"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var timeEntriesCollection *mongo.Collection = db.OpenCollection(db.Client, "timeEntries")

//Insert time entry into timeEntries collection
func InsertTimeEntry() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("1111")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var timeEntry models.TimeEntry
		var foundUser models.User
		if err := c.BindJSON(&timeEntry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithError(err).Error("Error binding JSON")
			return
		}
		log.Info(timeEntry.Tags)
		err := userCollection.FindOne(ctx, bson.M{"user_id": c.GetString("uid")}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.WithError(err).Error("Error finding user")
			return
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			log.WithError(err).Error("Error finding user2")
			return
		}

		timeEntry.ID = primitive.NewObjectID()
		timeEntry.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if timeEntry.Time_start.IsZero() == true {
			timeEntry.Time_start, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		}
		timeEntry.User_id = c.GetString("uid")

		resultInsertionNumber, insertErr := timeEntriesCollection.InsertOne(ctx, timeEntry)
		if insertErr != nil {
			msg := fmt.Sprintf("Time entry item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			log.WithError(insertErr).Error("Error inserting time entry")
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}
