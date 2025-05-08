FROM golang:1.24.0-alpine AS build

# Install dependencies to be able to build our code
# RUN apk add --no-cache --update git && apk add build-base

# Set the working directory inside the container
WORKDIR /app 

RUN apk add make

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest && go install github.com/swaggo/swag/cmd/swag@latest
