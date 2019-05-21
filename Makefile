GOOS?=linux
GOARCH?=amd64
GOCACHE?=/tmp/go-build
export GO111MODULE=on

gopath=$(shell echo ~)/go

export GOPATH=$(gopath)

.PHONY: clean
clean:
	rm -rf vendor
	rm -f server

.PHONY: build
build: clean dep
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o server .

.PHONY: dep
dep: source
	go mod vendor

.PHONY: source
source:
	@bash -c "source $(shell echo ~)/.bashrc"