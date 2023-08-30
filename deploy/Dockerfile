FROM golang:1.20 AS builder

ENV GO11MODULE=on \
    CGO_ENABLE=0 \
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

EXPOSE 8181

CMD ./app