package db

import (
	"context"
	"fmt"
	"time"

	"github.com/buker/go-api-starter/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	//"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection Retrieves a client to the MongoDB
func DBinstance() *mongo.Client {
	if err := config.Setup(); err != nil {
		log.WithError(err).Fatal("Failed to setup configuration")
	}
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")

	clusterEndpoint := viper.GetString("mongo.hostname")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	log.Info(connectionURI)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.WithError(err).Error("Failed to create client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to connect to cluster")
	}
	defer cancel()
	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.WithError(err).Error("Failed to ping cluster")
	}

	log.Info("Connected to MongoDB!")
	return client
}

//Client Database instance
var Client *mongo.Client = DBinstance()

//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbname := viper.GetString("mongo.dbname")
	var collection *mongo.Collection = client.Database(dbname).Collection(collectionName)

	return collection
}
