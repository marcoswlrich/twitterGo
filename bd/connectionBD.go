package bd

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/marcoswlrich/twittergo/models"
)

var (
	MongoCN      *mongo.Client
	DatabaseName string
)

func ConectarBD(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	passwd := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		user,
		passwd,
		host,
	)

	clientOptions := options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Coneção bem-sucedido com o banco de dados")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

func BaseConnected() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
