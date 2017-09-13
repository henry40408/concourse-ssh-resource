.PHONY: build down up

IMAGE_NAME=henry40408/alpine-ssh
CONTAINER_NAME=alpine-ssh

build:
	docker build -t $(IMAGE_NAME) .

up:
	docker run -d -p 22:22 --name $(CONTAINER_NAME) $(IMAGE_NAME)

down:
	docker stop -t 5 $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
