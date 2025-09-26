build:
	@go build -o bin/go-entrepreneurship-project-api ./cmd/go-entrepreneurship-project-api/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-entrepreneurship-project-api
