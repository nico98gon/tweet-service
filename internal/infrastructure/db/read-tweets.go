package db

import (
	"context"
	"tweet-service/internal/domain/tweets"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*tweets.ReturnTweets , bool) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")

	var results []*tweets.ReturnTweets
	condition := bson.M{"user_id": ID}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort((bson.D{{Key: "date", Value: -1}}))
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		return results, false
	}

	for cursor.Next(ctx) {
		var register tweets.ReturnTweets
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}

	return results, true
}