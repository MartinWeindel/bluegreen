REGISTRY              := eu.gcr.io/gardener-project
EXECUTABLE            := bluegreen
PROJECT               := github.com/MartinWeindel/bluegreen
IMAGE_REPOSITORY      := $(REGISTRY)/test/martinweindel/bluegreen
IMAGE_TAG             := "0.1"

.PHONY: build-local
build-local:
	@go build -o $(EXECUTABLE) \
	    -gcflags="all=-N -l" \
	    .

.PHONY: release
release:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(EXECUTABLE) \
	    -a \
	    -ldflags "-w" \
	    .

.PHONY: docker-images
docker-images:
	@docker build -t $(IMAGE_REPOSITORY):$(IMAGE_TAG) -f Dockerfile .
