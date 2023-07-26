package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/jwt"
	"github.com/marcoswlrich/twittergo/models"
)

func Login(ctx context.Context) models.RespApi {
	var t models.User
	var r models.RespApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Nome de usuario ou senha invalido " + err.Error()
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "O email do usuario e obrigatorio"
		return r
	}

	userData, existe := bd.TestandoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Nome de usuario ou senha invalido"
		return r
	}

	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Ocorreu um erro ao gerao o token correspondente > " + err.Error()
		return r
	}

	resp := models.RespostaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Ocorreu um erro ao tentar formatar o Token para Json"
	}

	cookies := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookiesString := cookies.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookiesString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
