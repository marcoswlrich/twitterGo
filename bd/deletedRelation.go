package bd

import (
	"context"

	"github.com/marcoswlrich/twittergo/models"
)

func DeletedRelation(t models.Relationship) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
