#build a small image
FROM busybox

USER root
WORKDIR /go/bin

COPY server /go/bin/server

CMD ["/go/bin/server"]
EXPOSE 1323