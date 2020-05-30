FROM golang:1.13.6-alpine3.11 as build
ARG POSTGRESQL_URL=""
# TODO: Fix App
ENV GO111MODULE=on \
    APP=/go/src/github.com/kyohei0423/sample/app \
    POSTGRESQL_URL=$POSTGRESQL_URL
WORKDIR $APP

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    bash \
    make \
    tzdata \
    gcc \
    musl-dev \
    git && \
    go get github.com/pilu/fresh && \
    go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" >  /etc/timezone && \
    apk del tzdata

COPY . $APP
