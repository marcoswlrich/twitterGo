package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func ListUsers(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUser := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Você deve enviar o parâmetro 'page' como um número inteiro maior que 0 " + err.Error()
		return r
	}

	users, status := bd.ReadAllUsers(IDUser, int64(pagTemp), search, typeUser)
	if !status {
		r.Message = "Erro ao ler todos os usuarios"
		return r
	}

	resJson, err := json.Marshal(users)
	if err != nil {
		r.Status = 500
		r.Message = "Erro ao formatar os dados dos usuários em JSON"
		return r
	}

	r.Status = 200
	r.Message = string(resJson)
	return r
}
