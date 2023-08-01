package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/marcoswlrich/twittergo/bd"
	"github.com/marcoswlrich/twittergo/models"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(
	ctx context.Context,
	uploadType string,
	request events.APIGatewayProxyRequest,
	claim models.Claim,
) models.RespApi {
	var r models.RespApi
	r.Status = 400
	IDuser := claim.ID.Hex()

	var filename string
	var user models.User

	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))
	switch uploadType {
	case "A":
		filename = "avatars/" + IDuser + ".jpg"
	case "B":
		filename = "banners/" + IDuser + ".jpg"
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		r.Status = 500
		r.Message = err.Error()
		return r
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = err.Error()
			return r
		}

		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			r.Status = 500
			r.Message = err.Error()
			return r
		}

		if err != io.EOF {
			if p.FileName() != "" {
				buf := bytes.NewBuffer(nil)
				if _, err := io.Copy(buf, p); err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}

				sess, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1"),
				})
				if err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}

				uploader := s3manager.NewUploader(sess)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(filename),
					Body:   &readSeeker{buf},
				})

				if err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}

			}
		}

		status, err := bd.EditRegister(user, IDuser)
		if err != nil || !status {
			r.Status = 400
			r.Message = "Erro ao editar o registro de usuario >" + err.Error()
			return r
		}

	} else {
		r.Message = "Deve-se enviar uma imagem com 'Context-Type' do tipo 'multipart/' no Header"
		r.Status = 400
		return r
	}

	r.Status = 200
	r.Message = "A imagem foi postada com sucesso"
	return r
}
