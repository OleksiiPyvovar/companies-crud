FROM golang:1.18.2-alpine3.16 AS builder


RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen

RUN apk update \
 && apk upgrade \
 && apk add ca-certificates gcc make libc-dev musl-dev binutils git curl jq \
 && rm -rf /var/cache/apk/*

ENV GOPATH=/go
ENV GOROOT=/usr/local/go

ENV APP_NAME=companies-crud
ENV APP_PKG=github.com/OleksiiPyvovar/${APP_NAME}
ENV APP_DIR=${GOPATH}/src/${APP_PKG}
ENV APP_BIN=/usr/local/bin/companies-crud

WORKDIR ${APP_DIR}

COPY api/ api/
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY go.mod go.mod
COPY go.sum go.sum


RUN go build -v -o ${APP_BIN} ./cmd/main.go

FROM alpine:3.16 as application

RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen

RUN apk update \
 && apk upgrade \
 && apk add curl jq mailcap tzdata ca-certificates \
 && rm -rf /var/cache/apk/*

ENV APP_BIN=/usr/local/bin/companies-crud

COPY --from=builder ${APP_BIN} ${APP_BIN}

CMD ${APP_BIN}
