package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/marcoswlrich/twittergo/awsgo"
	"github.com/marcoswlrich/twittergo/secretmanager"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(
	ctx context.Context,
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
