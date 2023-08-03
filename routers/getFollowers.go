package routers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func GetFollowers(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
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

	var resp models.GetQueryRelation

	okRelation := bd.QueryRelation(t)
	if !okRelation {
		resp.Status = false
	} else {
		resp.Status = true
	}
	respJson, err := json.Marshal(okRelation)
	if err != nil {
		r.Status = 500
		r.Message = "Erro ao formatar os dados dos usuarios como JSON " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
