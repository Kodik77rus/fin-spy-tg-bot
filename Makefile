.PHONY: build
.SILENT:

build:
	go mod download && CGO_ENABLED=0 go build -o ./bin ./cmd/fin-spy-tg-bot && ./bin/fin-spy-tg-bot -path ./configs/fin-spy-tg-bot.toml
