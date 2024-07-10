build: 
	@go build -o bin/pf *.go

run: build
	@./bin/pf

test:
	@go test -v ./...

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	protopkg/service.proto