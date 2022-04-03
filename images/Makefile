# Copyright 2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

include ../Makefile.defs

PUSH ?= false
ifeq ($(PUSH),true)
OUTPUT := "type=registry,push=true"
else
OUTPUT := "type=docker"
endif

#PLATFORM ?= linux/amd64,linux/arm64
PLATFORM ?= linux/amd64
IMAGE_REGISTRIES ?= docker.io/weizhoulan

.PHONY: all
all: controllerimage baseimage

.PHONY: controllerimage
controllerimage:
	@$(CONTAINER_ENGINE) buildx build -t $(IMAGE_REGISTRIES)/$@:$(VERSION) \
			--build-arg VERSION=$(VERSION) \
			--build-arg GIT_COMMIT_TIME=$(GIT_COMMIT_TIME) \
			--build-arg GIT_COMMIT_VERSION=$(GIT_COMMIT_VERSION) \
			--platform=$(PLATFORM) \
			--output $(OUTPUT)  -f $(ROOT_DIR)/images/controllerimage/Dockerfile $(ROOT_DIR)/

.PHONY: baseimage
baseimage:
	@$(CONTAINER_ENGINE) buildx build -t $(IMAGE_REGISTRIES)/$@:$(VERSION) \
		--platform=$(PLATFORM) --output $(OUTPUT) \
		-f $(ROOT_DIR)/images/controllerimage/Dockerfile $(ROOT_DIR)/

.PHONY: update-golang-image
update-golang-image:
	GO_VERSION=$(GO_VERSION) scripts/update-golang-image.sh


