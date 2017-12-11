# build stage

FROM golang:alpine as builder

ENV CGO_ENABLED 0

COPY . /go/src/github.com/henry40408/concourse-ssh-resource

RUN apk --no-cache add make && \
  cd /go/src/github.com/henry40408/concourse-ssh-resource && \
  make build-linux

WORKDIR /opt/resource

# release stage

FROM alpine:edge AS resource

RUN apk --no-cache add bash curl gzip jq tar openssl

COPY --from=builder /opt/resource /opt/resource
