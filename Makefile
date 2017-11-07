.PHONY: build build-linux clean test

GO_FLAGS := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
BUILD_ARGS := --ldflags "-s -w"

build:
	go build $(BUILD_ARGS) -o check ./cmd/check
	go build $(BUILD_ARGS) -o in ./cmd/in
	go build $(BUILD_ARGS) -o out ./cmd/out

build-linux:
	$(GO_FLAGS) go build $(BUILD_ARGS) -o /opt/resource/check ./cmd/check
	$(GO_FLAGS) go build $(BUILD_ARGS) -o /opt/resource/in ./cmd/in
	$(GO_FLAGS) go build $(BUILD_ARGS) -o /opt/resource/out ./cmd/out

clean:
	rm -f check in out

test:
	go test ./...
