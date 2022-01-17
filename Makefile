.DEFAULT_GOAL := all

all: install lint test build

install:
	@pnpm install && go mod vendor

lint:
	@pnpm run lint && gofmt -l ./main.go ./entity ./gateway ./port ./presenter ./repository ./usecase

format:
	@pnpm run format && gofmt -w ./main.go ./entity ./gateway ./port ./presenter ./repository ./usecase

test:
	@go test -cover -coverprofile=coverage.txt ./entity/... ./gateway/... ./port/... ./presenter/... ./repository/... ./usecase/...

build:
	@go build -mod vendor -o scrabbled
