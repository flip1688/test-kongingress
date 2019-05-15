
# STEP 1 build executable binary
FROM golang:alpine AS builder

RUN mkdir -p /go/src/github.com/flip1688/test-kongingress
WORKDIR /go/src/github.com/flip1688/test-kongingress

# Install dependencies.
RUN apk update && apk add --no-cache git

COPY go.mod ./
COPY go.sum ./

# Fetch dependencies.
RUN go mod download

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# STEP 2 build a small image
FROM busybox




USER root
WORKDIR /go/bin

COPY --from=builder /go/src/github.com/flip1688/test-kongingress/server /go/bin/server

CMD ["/go/bin/server"]
EXPOSE 1323