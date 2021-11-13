package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TimeEntry struct {
	ID         primitive.ObjectID `bson:"_id"`
	User_id    string             `json:"user_id"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Tags       *string            `json:"tags"`
	Time_start time.Time          `json:"time_start"`
	Time_end   time.Time          `json:"time_end"`
}
