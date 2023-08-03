package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/marcoswlrich/twittergo/models"
)

func QueryRelation(t models.Relationship) bool {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	condition := bson.M{
		"UserID":         t.UserID,
		"UserRelationID": t.UserRelationID,
	}

	var resultado models.Relationship
	err := col.FindOne(ctx, condition).Decode(&resultado)
	if err != nil {
		return false
	}

	return true
}
