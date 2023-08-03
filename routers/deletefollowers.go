package routers

import (
	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func DeleteFollowers(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "O parâmetro ID é obrigatório"
		return r
	}

	var t models.Relationship
	t.UserID = claim.ID.Hex()
	t.UserRelationID = ID

	status, err := bd.DeletedRelation(t)
	if err != nil {
		r.Message = "Ocorreu um erro ao excluir relation " + err.Error()
		return r
	}

	if !status {
		r.Message = "Falha ao excluir relacao"
	}

	r.Status = 200
	r.Message = "Excluido relacao com sucesso"
	return r
}
