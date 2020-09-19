# Tests to be run inside a docker container for clean environment.
# Tests to be added as code is written.
FROM golang:1.15.2-alpine

WORKDIR /app

COPY ./ /app

RUN go mod download && \
    go install zpltopng && \
    go test
