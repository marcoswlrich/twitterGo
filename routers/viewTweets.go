package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func ViewTweets(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	page := request.QueryStringParameters["page"]

	if len(ID) < 1 {
		r.Message = "O parametro ID e obrigatorio"
		return r
	}

	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Você deve enviar o parâmetro Page como um valor maior que 0"
		return r
	}

	tweets, correct := bd.VTweets(ID, int64(pag))
	if !correct {
		r.Message = "Erro ao ler os tweets"
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Erro ao formatar os dados do usuário como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
