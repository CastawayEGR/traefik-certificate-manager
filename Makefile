# Makefile Variables
CONTAINER_RUNTIME ?= podman
BIN_NAME := tcm-amd64-linux

build:
	rm -f go.mod go.sum
	@${CONTAINER_RUNTIME} build -t gobuild .
	@${CONTAINER_RUNTIME} container create --name temp gobuild
	@${CONTAINER_RUNTIME} container cp temp:/${BIN_NAME} .
	@${CONTAINER_RUNTIME} container cp temp:/go/src/tcm/go.mod .
	@${CONTAINER_RUNTIME} container cp temp:/go/src/tcm/go.sum .
	@${CONTAINER_RUNTIME} rm temp
	@${CONTAINER_RUNTIME} rmi localhost/gobuild


clean:
	@${CONTAINER_RUNTIME} image prune -f
	@${CONTAINER_RUNTIME} rmi $(@${CONTAINER_RUNTIME} image list --filter dangling=true -q) -f
