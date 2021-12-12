IMAGE=docker.io/biancarosa/shutting-down-gracefully
IMAGE_TAG=$(shell git rev-parse --short=7 HEAD)

build:
	podman build . -t $(IMAGE):$(IMAGE_TAG)

push:
	podman push $(IMAGE):$(IMAGE_TAG)