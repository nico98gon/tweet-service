package tweets

import (
	"time"
)

type CreateTweet struct {
	Content 	string 							`bson:"content" json:"content,omitempty"`
	UserID    string	`bson:"userid" json:"user,omitempty"`
	Date 			time.Time 					`bson:"date" json:"date,omitempty"`
}