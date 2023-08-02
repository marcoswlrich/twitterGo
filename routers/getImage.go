package routers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/marcoswlrich/twittergo/awsgo"
	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

func GetImage(
	ctx context.Context,
	uploadType string,
	request events.APIGatewayProxyRequest,
	claim models.Claim,
) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "O parâmetro ID é obrigatório"
		return r
	}

	perfil, err := bd.BuscandoPerfil(ID)
	if err != nil {
		r.Message = "Usuario nao encontrado " + err.Error()
		return r
	}

	var filename string
	switch uploadType {
	case "A":
		filename = perfil.Avatar
	case "B":
		filename = perfil.Banner
	}

	fmt.Println("Filename " + filename)
	svc := s3.NewFromConfig(awsgo.Cfg)

	file, err := downloadFromS3(ctx, svc, filename)
	if err != nil {
		r.Status = 500
		r.Message = "Erro ao baixar arquivo da S3 " + err.Error()
		return r
	}

	r.CustomResp = &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       file.String(),
		Headers: map[string]string{
			"Context-Type":        "application/octet-stream",
			"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
		},
	}
	return r
}

func downloadFromS3(ctx context.Context, svc *s3.Client, filename string) (*bytes.Buffer, error) {
	bucket := ctx.Value(models.Key("bucketName")).(string)
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}

	defer obj.Body.Close()
	fmt.Println("bucketname = " + bucket)

	file, err := ioutil.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(file)

	return buffer, nil
}
