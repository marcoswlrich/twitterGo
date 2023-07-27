package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/jwt"
	"github.com/marcoswlrich/twittergo/models"
	"github.com/marcoswlrich/twittergo/routers"
)

func Manipuladores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println(
		"Processando " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string),
	)

	var r models.RespApi
	r.Status = 400

	isOk, statuscode, msg, claim := validaAuthorization(ctx, request)
	if !isOk {
		r.Status = statuscode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		case "login":
			return routers.Login(ctx)
		}

	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "perfil":
			return routers.ViewPerfil(request)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "editarPerfil":
			return routers.EditPerfil(ctx, claim)
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
	}

	r.Message = "Method Invalid"
	return r
}

func validaAuthorization(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "registro" || path == "login" || path == "obterAvatar" || path == "obterBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token obrigatorio", models.Claim{}
	}

	claim, tudoOk, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !tudoOk {
		if err != nil {
			fmt.Println("Error token" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Erro Token" + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token Ok")
	return true, 200, msg, *claim
}
