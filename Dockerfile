#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/fin-spy-tg-bot

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/bin /fin-spy-tg-bot
USER nonroot:nonroot
ENTRYPOINT [ "/fin-spy-tg-bot" ]sudo apt-get install apparmor
LABEL Name=finspytgbot Version=0.0.1

