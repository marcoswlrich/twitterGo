package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

var (
	Email  string
	IDUser string
)

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de Token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		// Verificacao de rotins no BD
		_, encontrado, _ := bd.ChecarSeExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return &claims, encontrado, IDUser, nil
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token invalido")
	}

	return &claims, false, string(""), errors.New("Token invalido")
}
