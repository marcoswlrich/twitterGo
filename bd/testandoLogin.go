package bd

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/marcoswlrich/twittergo/models"
)

func TestandoLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := ChecarSeExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true
}
