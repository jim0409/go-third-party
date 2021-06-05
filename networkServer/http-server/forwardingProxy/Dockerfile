FROM golang:1.11.5-alpine3.8 AS builder

# basic packages needed
# RUN set -eux; \
        # apk update && \
        # apk add --no-cache --virtual .build-deps gcc libc-dev git

# env for go workdir and mod vendor
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/github.com/jimweng/forwardingproxy
# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR
# copy files into workdir
ADD . $GO_WORKDIR
RUN go build -o whoyare -i whoyare.go

FROM alpine:3.8
RUN set -eux; \
        apk add --no-cache ca-certificates wget openssl
# copy binary into container
COPY --from=builder /go/src/github.com/jimweng/forwardingproxy/whoyare .
# ADD ./entrypoint.sh .
# RUN chmod +x entrypoint.sh

# CMD ["./entrypoint.sh"]
CMD ["./whoyare"]