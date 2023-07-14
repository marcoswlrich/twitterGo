package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/models"
)

func Manipuladores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println(
		"Processando " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string),
	)

	var r models.RespApi
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		}

	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
	}

	r.Message = "Method Invalid"
	return r
}
