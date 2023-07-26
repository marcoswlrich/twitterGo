package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func ViewPerfil(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entrei no viewProfile")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "O parametro ID e obrigatorio"
		return r
	}

	perfil, err := bd.BuscandoPerfil(ID)
	if err != nil {
		r.Message = "Ocorreu um erro ao tentar buscar o registro > " + err.Error()
		return r
	}

	resJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Erro ao formatar os dados do usuario como JSON > " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(resJson)
	return r
}
