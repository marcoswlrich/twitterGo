package routers

import (
	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "O parametro ID e obrigatorio"
		return r
	}

	err := bd.DeletedTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocorreu um erro ao tentar excluir o tweet >" + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Tweet excluido com sucesso"
	return r
}
