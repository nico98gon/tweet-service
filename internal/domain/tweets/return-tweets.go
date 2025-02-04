package tweets

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnTweets struct {
	ID 			primitive.ObjectID	`bson:"_id" json:"_id,omitempty"`
	UserID  string							`bson:"user_id" json:"user_id,omitempty"`
	Content	string							`bson:"content" json:"content,omitempty"`
	Date 		time.Time 					`bson:"date" json:"date,omitempty"`
}