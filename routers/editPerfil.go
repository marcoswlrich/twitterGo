package routers

import (
	"context"
	"encoding/json"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func EditPerfil(ctx context.Context, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	var t models.User

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Dados incorretos" + err.Error()
	}

	status, err := bd.EditRegister(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocorreu um erro ao tentar atualizar o registro." + err.Error()
		return r
	}

	if !status {
		r.Message = "Falha ao atualizar registro do usuario."
		return r
	}

	r.Status = 200
	r.Message = "Perfil atualizado com sucesso"
	return r
}
