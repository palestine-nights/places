VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo latest)
SERVICE ?= places
IMAGE   := ashanaakh/$(SERVICE):$(VERSION)

.PHONY: default build push run

default: build run

run:
	@echo '> Starting "$(SERVICE)" container...'
	@docker run --name $(SERVICE) -p 8000:8000 -d $(IMAGE)

push: build
	@echo '> Pushing "$(SERVICE)" docker image with version: "$(VERSION)"'
	@docker push $(IMAGE)

build:
	@echo '> Building "$(SERVICE)" docker image...'
	@docker build -t $(IMAGE) .

clean:
	@docker rm -f $(SERVICE)