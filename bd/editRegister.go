package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/marcoswlrich/twittergo/models"
)

func EditRegister(u models.User, ID string) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	registro := make(map[string]interface{})
	if len(u.Name) > 0 {
		registro["name"] = u.Name
	}
	if len(u.Apelidos) > 0 {
		registro["apelidos"] = u.Apelidos
	}
	registro["dataNascimento"] = u.DataNascimento
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Publication) > 0 {
		registro["publication"] = u.Publication
	}
	if len(u.SiteWeb) > 0 {
		registro["siteweb"] = u.SiteWeb
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
