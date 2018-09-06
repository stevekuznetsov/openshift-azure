# all is the default target to build everything
all: clean build sync

build:
	go build ./...

clean:
	rm -f sync

test:
	go test ./...

generate:
	go generate ./...

logbridge: clean generate
	go build ./cmd/logbridge

logbridge-image: logbridge
	go get github.com/openshift/imagebuilder/cmd/imagebuilder
	imagebuilder -f Dockerfile.logbridge -t quay.io/openshift-on-azure/logbridge:latest .

logbridge-push: logbridge-image
	docker push quay.io/openshift-on-azure/logbridge:latest

sync: clean generate
	CGO_ENABLED=0 go build ./cmd/sync

TAG ?= $(shell git rev-parse --short HEAD)
SYNC_IMAGE ?= quay.io/openshift-on-azure/sync:$(TAG)

sync-image: sync
	go get github.com/openshift/imagebuilder/cmd/imagebuilder
	imagebuilder -f Dockerfile.sync -t $(SYNC_IMAGE) .

sync-push: sync-image
	docker push $(SYNC_IMAGE)

verify:
	./hack/validate-generated.sh

unit:
	go test ./...

.PHONY: clean sync-image sync-push verify unit
