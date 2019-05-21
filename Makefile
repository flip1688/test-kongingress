export GOOS=linux
export GOARCH=amd64
export GOCACHE=/tmp/go-build
export GO111MODULE=on
.PHONY: build
build: dep
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o server .

.PHONY: dep
dep:
	@go mod vendor