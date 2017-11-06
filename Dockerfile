FROM golang:1.9-alpine

COPY . /go/src/github.com/henry40408/concourse-ssh-resource

RUN apk --no-cache add make && \
    cd /go/src/github.com/henry40408/concourse-ssh-resource && \
    make build-linux && \
    rm -r /go && \
    apk del make

WORKDIR /opt/resource
