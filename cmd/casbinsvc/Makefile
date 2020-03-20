build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build .

docker:
	docker build -t casbinsvc ./

clear:
	rm ./casbinsvc