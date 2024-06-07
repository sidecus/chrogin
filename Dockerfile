# syntax=docker/dockerfile:1.4

# Release stage, use TINI to collect zombie processes after executing chrominum
FROM alpine:3.20.0 as base
ARG APKREPOSITORY
RUN if [ -n "${APKREPOSITORY}" ]; then \
        sed -i "s/dl-cdn.alpinelinux.org/${APKREPOSITORY}/g" /etc/apk/repositories; \
    fi
RUN apk add --update --no-cache chromium tini

# Build stage, use custom mavensettings to setup mirror
FROM golang:1.22.3-alpine3.20 as builder
ARG GOPROXY
ENV GOPROXY=${GOPROXY}
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o ./publish/ .

# dev container only
FROM base as devcontainer
RUN apk add --update --no-cache git
ENV GOPATH=/go \
    GOROOT=/usr/local/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"
COPY --from=golang:1.22.3-alpine3.20 --link $GOROOT $GOROOT
COPY --chmod=0755 ./.vscode/gotools.sh $GOPATH/
RUN sh -c $GOPATH/gotools.sh

# Release stage, use TINI to collect zombie processes after executing chrominum
FROM base as release
WORKDIR /app
COPY --from=builder --chmod=0555 /app/publish/chrogin ./
COPY templates templates/
COPY assets assets/
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/chrogin"]
