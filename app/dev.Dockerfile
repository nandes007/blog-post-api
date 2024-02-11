# syntax=docker/dockerfile:1

ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine

ENV GO_ENV production

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 9001

# Run
CMD ["air", "-c", ".air.toml"]