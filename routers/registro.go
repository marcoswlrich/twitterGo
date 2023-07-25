package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.User
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entrar em registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "E preciso adicionar o email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "E preciso adicionar uma senha de pelo menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChecarSeExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ja existe um usuario registrado com esse email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.AdicionarRegistro(t)
	if err != nil {
		r.Message = "Ocorreu um erro ao tentar realizar o registro de usuario" + err.Error()
		return r
	}

	if !status {
		r.Message = "Falha ao inserir registro do usuário "
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registrado com sucesso"
	fmt.Println(r.Message)
	return r
}
