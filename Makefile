.PHONY: build
.SILENT:

build:
	go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/fin-spy-tg-bot && ./bin
