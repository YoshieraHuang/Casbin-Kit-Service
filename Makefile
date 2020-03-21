# build:
# 	GOOS=linux CGO_ENABLE=0 GOARCH=amd64 go build ./cmd/casbinsvc

docker:
	docker build -t casbinsvc ./

clear:
	rm ./cmd/casbinsvc/casbinsvc