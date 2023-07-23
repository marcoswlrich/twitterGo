package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	custo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), custo)
	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil
}
