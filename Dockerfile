FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY ["go.mod", "go.sum", "./" ]

RUN go mod download

COPY . *.go ./

RUN go build -o ./bot ./cmd/fin-spy-tg-bot

##
## Deploy
##

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /app/bot /bot

ENTRYPOINT ["/bot"]

LABEL Name=tg-bot Version=0.0.1
