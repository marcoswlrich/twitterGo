package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func ReadTweetsFollowers(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400
	IDUser := claim.ID.Hex()

	page := request.QueryStringParameters["page"]
	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Você deve enviar o parâmetro Page como um valor maior que 0"
		return r
	}

	tweets, correct := bd.ReadTweetsFoll(IDUser, pag)
	if !correct {
		r.Message = "Erro ao ler os tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Erro ao formatar os dados dos tweets dos seguidores"
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
