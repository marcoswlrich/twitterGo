export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
.DEFAULT_GOAL := deploy

deploy:
	go build -o main
	zip -r function.zip main
