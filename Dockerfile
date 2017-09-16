FROM golang:1.9-alpine AS builder

COPY . /go/src/github.com/henry40408/ssh-shell-resource
COPY Makefile /go/Makefile

RUN apk add --update make && \
    make build_in_docker

FROM alpine

COPY --from=builder /assets /opt/resource
