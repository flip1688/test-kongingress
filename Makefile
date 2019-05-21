
export GO111MODULE=on
.PHONY: build
build: dep
	CGO_ENABLED=0 go build -o server .

.PHONY: dep
dep:
	@go mod vendor