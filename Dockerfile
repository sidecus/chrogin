# Build stage, use custom mavensettings to setup mirror
FROM golang:1.22.3-alpine3.20 as builder
ARG GOPROXY
ENV GOPROXY=${GOPROXY}
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o ./publish/ .

# Release stage, use TINI to collect zombie processes after executing chrominum
FROM alpine:3.20.0
ARG APKREPOSITORY
RUN if [ -n "${APKREPOSITORY}" ]; then \
        sed -i "s/dl-cdn.alpinelinux.org/${APKREPOSITORY}/g" /etc/apk/repositories; \
    fi
RUN apk add --update --no-cache chromium tini
WORKDIR /app
COPY --from=builder --chmod=0555 /app/publish/chrogin ./
COPY templates templates/
COPY assets assets/
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/chrogin"]
