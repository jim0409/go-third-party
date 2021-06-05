FROM golang:1.11.5-alpine3.8 AS builder

# basic packages needed
RUN set -eux; \
        apk update && \
        apk add --no-cache --virtual .build-deps gcc libc-dev git

# env for go workdir and mod vendor
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/github.com/jimweng/crawler/crawler
ENV GO111MODULE=on
# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR
# copy files into workdir
ADD . $GO_WORKDIR
RUN go mod vendor
RUN cd cmd/src; go build -o crawler -i main.go

FROM alpine:3.8
RUN set -eux; \
        apk add --no-cache ca-certificates wget openssl
# copy binary into container
COPY --from=builder /go/src/github.com/jimweng/crawler/crawler/cmd/src .
ADD ./entrypoint.sh .
RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]