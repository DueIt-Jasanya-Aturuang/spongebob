FROM golang:1.20 AS builder

RUN mkdir /app
RUN cd /app

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