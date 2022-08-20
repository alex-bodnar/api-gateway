ARG GO_VERSION=1.18.3
ARG ALPINE_VERSION=3.16

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

ARG APP_NAME
ARG VERSION
ARG COMMIT_DATE
ARG FORTUNE_COOKIE
ARG TAG
ARG COMMIT_LAST

WORKDIR /build
COPY . /build/

RUN go build --tags muslc -mod=vendor -ldflags=" \
    -X main.appName=$APP_NAME \
    -X main.date=$COMMIT_DATE \
    -X main.fortuneCookie=$FORTUNE_COOKIE \
    -X main.version=$VERSION \
    -X main.tag=$TAG \
    -X main.commit=$COMMIT_LAST" \
    -o /build/$APP_NAME /build/cmd/main.go

FROM alpine:${ALPINE_VERSION}
ARG APP_NAME
ENV BINARY_NAME=$APP_NAME
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /build/$APP_NAME /usr/bin/$APP_NAME
ENTRYPOINT /usr/bin/$BINARY_NAME -c config.yaml
