build:
	@go build -o bin/project-builder-api ./cmd/project-builder-api/main.go

test:
	@go test -v ./...

run: build
	@./bin/project-builder-api
