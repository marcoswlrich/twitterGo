package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func CreateTweet(ctx context.Context, claim models.Claim) models.RespApi {
	var mensagem models.Tweet
	var r models.RespApi
	r.Status = 400
	IDUser := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensagem)
	if err != nil {
		r.Message = "Ocorreu um erro ao tentar decodificar o corpo " + err.Error()
		return r
	}

	registro := models.SaveTweet{
		UserID:   IDUser,
		Mensagem: mensagem.Mensagem,
		Data:     time.Now(),
	}

	_, status, err := bd.InsertTweet(registro)
	if err != nil {
		r.Message = "Ocorreu um erro ao tentar inserir o registro " + err.Error()
		return r
	}

	if !status {
		r.Message = "Falha ao tentar inserir o Tweet"
		return r
	}

	r.Status = 200
	r.Message = "Tweet criado com sucesso"
	return r
}
