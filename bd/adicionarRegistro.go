package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/marcoswlrich/twittergo/models"
)

func AdicionarRegistro(u models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
