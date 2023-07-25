package jwt

import (
	"context"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/marcoswlrich/twittergo/models"
)

func GeneroJWT(ctx context.Context, t models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	minhaChave := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":           t.Email,
		"name":            t.Name,
		"apelido":         t.Apelidos,
		"data_nascimento": t.DataNascimento,
		"biografia":       t.Biografia,
		"publicacao":      t.Publication,
		"siteWeb":         t.SiteWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(minhaChave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
