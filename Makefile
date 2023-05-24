SERVICE_NAME=go-simple-api
GIT_COMMIT_HASH=$(shell git rev-parse HEAD | cut -c -16)

.PHONY: all build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./${SERVICE_NAME} ./main.go
	docker build -t lisyaoran51/${SERVICE_NAME}:${GIT_COMMIT_HASH} . 
	docker tag lisyaoran51/${SERVICE_NAME}:${GIT_COMMIT_HASH} lisyaoran51/${SERVICE_NAME}
	docker push lisyaoran51/${SERVICE_NAME}