package routers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func AddFollowers(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
	claim models.Claim,
) models.RespApi {
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

	status, err := bd.InsertRelation(t)
	if err != nil {
		r.Message = "Ocorreu um erro ao adicionar relation " + err.Error()
		return r
	}

	if !status {
		r.Message = "Falha ao inserir relacao"
	}

	r.Status = 200
	r.Message = "Add relacao com sucesso"
	return r
}
