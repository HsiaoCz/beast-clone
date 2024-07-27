run: build
	@./bin/sm

build:
	@go build -o bin/sm main.go

test:
	@go test -v ./...

mongo:
	@docker start mongo
