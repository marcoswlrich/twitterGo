package bd

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/marcoswlrich/twittergo/models"
)

func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	var results []*models.User

	options := options.Find()
	options.SetLimit(20)
	options.SetSkip((page - 1) * 20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, options)
	if err != nil {
		return results, false
	}

	var include bool

	for cur.Next(ctx) {
		var s models.User

		err := cur.Decode(&s)
		if err != nil {
			fmt.Println("Decode = " + err.Error())
			return results, false
		}

		var r models.Relationship
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found := QueryRelation(r)
		if tipo == "new" && !found {
			include = true
		}
		if tipo == "follow" && found {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include {
			s.Password = ""
			results = append(results, &s)
		}

	}

	err = cur.Err()
	if err != nil {
		fmt.Println("cur.Err() = " + err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
