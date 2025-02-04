package db

import (
	"context"
	"tweet-service/internal/domain/tweets"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t tweets.CreateTweet) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")

	register := bson.M{
		"user_id":  t.UserID,
		"content":  t.Content,
		"date":     t.Date,
		// "likes":    0,
		// "retweets": 0,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}