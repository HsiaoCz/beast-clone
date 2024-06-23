run: build
	@./bin/tw

build:
	@go build -o bin/tw main.go

test:
	@go test -v ./...

gen:
	@templ generate

mongo:
	@docker start mongo