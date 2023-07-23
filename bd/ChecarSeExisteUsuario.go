package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/marcoswlrich/twittergo/models"
)

func ChecarSeExisteUsuario(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
