FROM golang:1.20.2 as builder

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN go mod tidy
