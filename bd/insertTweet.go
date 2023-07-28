package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/marcoswlrich/twittergo/models"
)

func InsertTweet(t models.SaveTweet) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":   t.UserID,
		"mensagem": t.Mensagem,
		"data":     t.Data,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
