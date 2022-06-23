build:
	cd cmd && go build -o /go/bin/server 

run:
	go run cmd/main.go

install:
	go get -d -v ./...
	go install -v ./...

test:
	go test ./...

format:
	gofmt -w -s .

docker:
	docker build -t gitlab_project .