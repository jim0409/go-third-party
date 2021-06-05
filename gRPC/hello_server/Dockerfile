FROM golang:1.11.5-alpine3.8 AS builder

# basic packages needed
RUN set -eux; \
        apk update && \
        apk add --no-cache --virtual .build-deps gcc libc-dev git

# env for go workdir and mod vendor
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/go-third-party/gRPC/hello_server/
# ENV GO111MODULE=on
# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR
# copy files into workdir
ADD . $GO_WORKDIR

WORKDIR $GO_WORKDIR/server
RUN go get
RUN go build .

CMD ["./server"]