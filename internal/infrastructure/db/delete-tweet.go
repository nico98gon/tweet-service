package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, userID string) (error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
		"user_id": userID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	if err != nil {
		return err
	}

	return nil
}