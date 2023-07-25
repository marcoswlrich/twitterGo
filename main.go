package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/marcoswlrich/twittergo/awsgo"
	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/handlers"
	"github.com/marcoswlrich/twittergo/models"
	"github.com/marcoswlrich/twittergo/secretmanager"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(
	request events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InitAWS()

	if !ValidateParameters() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error variaveis de retorno, deve incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Contert-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error na leitura de Segredo" + err.Error(),
			Headers: map[string]string{
				"Contert-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// Conexao com banco de dados

	err = bd.ConectarBD(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error ao conectar banco de dados" + err.Error(),
			Headers: map[string]string{
				"Contert-Type": "application/json",
			},
		}
		return res, nil

	}

	respAPI := handlers.Manipuladores(awsgo.Ctx, request)
	if respAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Contert-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return respAPI.CustomResp, nil
	}
}

func ValidateParameters() bool {
	_, bringParameter := os.LookupEnv("SecretName")
	if !bringParameter {
		return bringParameter
	}

	_, bringParameter = os.LookupEnv("BucketName")
	if !bringParameter {
		return bringParameter
	}

	_, bringParameter = os.LookupEnv("UrlPrefix")
	if !bringParameter {
		return bringParameter
	}

	return bringParameter
}
