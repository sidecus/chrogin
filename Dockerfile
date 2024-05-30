ARG GOLANG_VERSION=1.22.3

# Build stage, use custom mavensettings to setup mirror
FROM golang:${GOLANG_VERSION}-alpine3.20 as builder
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
COPY ./go.* .
RUN go mod tidy
COPY . .
RUN go build -o ./publish/ .

# Release stage, use TINI to collect zombie processes after executing chrominum
FROM alpine:3.20.0
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add --update --no-cache chromium tini

# USER user
WORKDIR /app
COPY --from=builder --chmod=0555 /app/publish/gin-report /app
COPY ./templates /app/templates
COPY ./assets /app/assets
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/gin-report"]
