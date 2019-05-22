#build a small image
FROM scratch

WORKDIR /go/bin

COPY server /go/bin/server

CMD ["/go/bin/server"]
EXPOSE 1323