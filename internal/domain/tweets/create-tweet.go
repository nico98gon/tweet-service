package tweets

import (
	"time"
)

type CreateTweet struct {
	Content 	string 		`bson:"content" json:"content,omitempty"`
	UserID    string		`bson:"user_id" json:"user_id,omitempty"`
	Date 			time.Time	`bson:"date" 		json:"date,omitempty"`
}