package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/marcoswlrich/twittergo/models"
)

func VTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweet")

	var resultados []*models.ReturnTweets

	condicion := bson.M{
		"userid": ID,
	}

	opcoes := options.Find()
	opcoes.SetLimit(20)
	opcoes.SetSort(bson.D{{Key: "data", Value: -1}})
	opcoes.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opcoes)
	if err != nil {
		return resultados, false
	}

	for cursor.Next(ctx) {
		var registro models.ReturnTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
