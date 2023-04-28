FROM --platform=linux/amd64 golang:1.20.3-alpine3.17 AS builder

# basic packages needed
RUN set -eux; \
        apk update && \
        apk add --no-cache --virtual .build-deps gcc libc-dev git

# env for go workdir and mod vendor
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/etherscan
ENV GO111MODULE=on
# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR
# copy files into workdir
ADD . $GO_WORKDIR
RUN go mod vendor
RUN go build .

FROM --platform=linux/amd64 alpine:3.17
RUN set -eux; \
        apk add --no-cache ca-certificates wget openssl

# copy binary into container
COPY --from=builder /go/src/etherscan .
RUN ["./etherscan"]
