FROM golang:1.22.1 AS builder
WORKDIR /app
ARG GITHUB_REPO=shamrin/perusnetti
RUN git clone https://github.com/${GITHUB_REPO} .
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o perusnetti .
