FROM golang:1.9-alpine

COPY . /go/src/github.com/henry40408/concourse-ssh-resource

RUN apk add --update git make && \
    go get -u github.com/kardianos/govendor && \
    cd /go/src/github.com/henry40408/concourse-ssh-resource && \
    govendor sync && \
    make build-linux && \
    cd /opt/resource && \
    rm -r /go && \
    apk del --purge git make

WORKDIR /opt/resource
