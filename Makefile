build:
	cd cmd && go build -o /go/bin/server 

run:
	go run cmd/main.go

test:
	go test ./...

format:
	gofmt -w -s .

docker:
	docker build -t gitlab_project .