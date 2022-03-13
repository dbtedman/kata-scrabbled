.DEFAULT_GOAL := all

all: install lint test build

install:
	@pnpm install && go mod vendor

lint:
	@pnpm run lint && gofmt -l ./main.go ./cmd ./internal ./web

format:
	@pnpm run format && gofmt -w ./main.go ./cmd ./internal ./web

test:
	@go test -race -cover -coverprofile=coverage.txt ./cmd/.. ./internal/... ./web/...

build:
	@go build -race -mod vendor -o scrabbled ./cmd/scrabbled
