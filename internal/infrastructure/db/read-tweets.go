package db

import (
	"context"
	"tweet-service/internal/domain/tweets"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, cursor string) ([]*tweets.ReturnTweets, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")

	var results []*tweets.ReturnTweets

	filter := bson.M{"user_id": ID}
	if cursor != "" {
		filter["_id"] = bson.M{"$lt": cursor}
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})

	cursorResult, err := col.Find(ctx, filter, options)
	if err != nil {
		return results, false, err
	}

	for cursorResult.Next(ctx) {
		var register tweets.ReturnTweets
		err := cursorResult.Decode(&register)
		if err != nil {
			return results, false, err
		}
		results = append(results, &register)
	}

	hasMore := len(results) == 20

	return results, hasMore, nil
}