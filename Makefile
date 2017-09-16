.PHONY: build build_in_docker

BUILD_ARGS=--ldflags "-s -w"

build:
	go build $(BUILD_ARGS) -o check cmd/check
	go build $(BUILD_ARGS) -o in cmd/in
	go build $(BUILD_ARGS) -o out cmd/out

build_in_docker:
	go build $(BUILD_ARGS) -o /assets/check github.com/henry40408/ssh-shell-resource/cmd/check
	go build $(BUILD_ARGS) -o /assets/in github.com/henry40408/ssh-shell-resource/cmd/in
	go build $(BUILD_ARGS) -o /assets/out github.com/henry40408/ssh-shell-resource/cmd/out
