package db

import (
	"context"
	"time"

	"github.com/buker/go-api-starter/internal/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimeEntry struct{}

var collection = new(mongo.Collection)

func (p *TimeEntry) Insert(timeEntry models.TimeEntry) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, &timeEntry)
	log.Debug("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err
}
func (p *TimeEntry) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	if err != nil {
		log.WithError(err).Fatal()
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	log.Debug("Deleted a single document: ", result.DeletedCount)
	return err
}

// Get all Places
func (p *TimeEntry) FindAll() ([]models.TimeEntry, error) {
	var timeEntries []models.TimeEntry

	findOptions := options.Find()
	findOptions.SetLimit(100)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// Finding multiple documents returns a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.TimeEntry
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		timeEntries = append(timeEntries, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return timeEntries, err
}
