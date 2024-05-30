# Build stage, use custom mavensettings to setup mirror
FROM golang:1.22.3-alpine3.20 as builder
ARG GOPROXY
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=${GOPROXY}
COPY ./go.* .
RUN go mod download
COPY ./*.go .
RUN go build -o ./publish/ .

# Release stage, use TINI to collect zombie processes after executing chrominum
FROM alpine:3.20.0
ARG APKREPOSITORY=""
RUN if [ -n "${APKREPOSITORY}" ]; then \
        sed -i "s/dl-cdn.alpinelinux.org/${APKREPOSITORY}/g" /etc/apk/repositories; \
    fi
RUN apk add --update --no-cache chromium tini
WORKDIR /app
COPY --from=builder --chmod=0555 /app/publish/chrogin /app
COPY ./templates /app/templates
COPY ./assets /app/assets
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/chrogin"]
