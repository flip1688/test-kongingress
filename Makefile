export GOOS=linux
export GOARCH=amd64
export GOCACHE=/tmp/go-build
export GO111MODULE=on

gopath=$(shell echo ~)/go

export GOPATH=$(gopath)

.PHONY: build
build: dep
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o server .

.PHONY: dep
dep: source
	@go mod vendor

.PHONY: source
source:
	@source $(shell echo ~)/.bashrc
