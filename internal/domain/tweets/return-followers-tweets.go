package tweets

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type  ReturnFollowersTweets struct {
	ID 				primitive.ObjectID	`bson:"_id" json:"_id,omitempty"`
	UserID  	string							`bson:"user_id" json:"user_id,omitempty"`
	UserIDRel string							`bson:"user_id_rel" json:"user_id_rel,omitempty"`
	Tweet			tweet
}

type tweet struct {
	Content string 		`bson:"content" json:"content,omitempty"`
	Date 		time.Time `bson:"date" json:"date,omitempty"`
	ID 			string		`bson:"_id" json:"_id,omitempty"`
}