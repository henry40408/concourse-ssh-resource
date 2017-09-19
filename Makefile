.PHONY: build build_in_docker test

GO_FLAGS := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
BUILD_ARGS := --ldflags "-s -w"

build:
	go build $(BUILD_ARGS) -o check cmd/check
	go build $(BUILD_ARGS) -o in cmd/in
	go build $(BUILD_ARGS) -o out cmd/out

build_in_docker:
	$(GO_FLAGS) go build $(BUILD_ARGS) \
		-o /opt/resource/check \
		github.com/henry40408/ssh-shell-resource/cmd/check
	$(GO_FLAGS) go build $(BUILD_ARGS) \
		-o /opt/resource/in \
		github.com/henry40408/ssh-shell-resource/cmd/in
	$(GO_FLAGS) go build $(BUILD_ARGS) \
		-o /opt/resource/out \
		github.com/henry40408/ssh-shell-resource/cmd/out

test:
	go test github.com/henry40408/ssh-shell-resource/cmd/check
	go test github.com/henry40408/ssh-shell-resource/cmd/in
	go test github.com/henry40408/ssh-shell-resource/cmd/out
