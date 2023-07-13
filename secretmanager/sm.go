package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/marcoswlrich/twittergo/awsgo"
	"github.com/marcoswlrich/twittergo/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dadosSecretos models.Secret
	fmt.Println("> Pedido secreto" + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dadosSecretos, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &dadosSecretos)
	fmt.Println("> Leitura de segredo Ok" + secretName)

	return dadosSecretos, nil
}
