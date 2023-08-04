package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/marcoswlrich/twittergo/models"
)

func ReadTweetsFoll(ID string, page int) ([]models.ReturnTweetsFollowers, bool) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	skip := (page - 1) * 20

	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{"$match": bson.M{"userid": ID}})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condition = append(condition, bson.M{"$unwind": "$tweet"})
	condition = append(condition, bson.M{"$sort": bson.M{"tweet.data": -1}})
	condition = append(condition, bson.M{"$skip": skip})
	condition = append(condition, bson.M{"$limit": 20})

	var result []models.ReturnTweetsFollowers

	cursor, err := col.Aggregate(ctx, condition)
	if err != nil {
		return result, false
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
