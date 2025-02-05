package db

import (
	"context"
	"tweet-service/internal/domain/tweets"

	"go.mongodb.org/mongo-driver/bson"
)

func GetFollowersTweets(ID string, cursor string) ([]tweets.ReturnFollowersTweets, string, bool) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	matchStage := bson.M{"$match": bson.M{"user_id": ID}}
	lookupStage := bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "user_id_rel",
			"foreignField": "user_id",
			"as":           "tweet",
		},
	}
	unwindStage := bson.M{"$unwind": "$tweet"}
	sortStage := bson.M{"$sort": bson.M{"tweet.date": -1}}
	limitStage := bson.M{"$limit": 20}

	if cursor != "" {
		matchStage["$match"] = bson.M{
			"user_id": ID,
			"tweet._id": bson.M{"$lt": cursor},
		}
	}

	pipeline := []bson.M{matchStage, lookupStage, unwindStage, sortStage, limitStage}

	var result []tweets.ReturnFollowersTweets

	cursorResult, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return result, "", false
	}

	err = cursorResult.All(ctx, &result)
	if err != nil {
		return result, "", false
	}

	var nextCursor string
	if len(result) > 0 {
		nextCursor = result[len(result)-1].Tweet.ID
	}

	return result, nextCursor, true
}