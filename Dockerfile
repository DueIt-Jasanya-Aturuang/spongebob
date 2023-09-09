FROM golang:1.20 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o account .

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/account /app

WORKDIR /app

EXPOSE 7002

CMD ./account