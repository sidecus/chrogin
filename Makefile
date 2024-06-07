NAME := chrogin
TAG := $(shell git log -1 --pretty=%h)
NAMESPACE := sidecus

.PHONY: debug build buildcn full run

debug: buildcn run

build:
	docker build -t ${NAMESPACE}/${NAME}:latest .
	# docker build --progress=plain -t ${NAMESPACE}/${NAME}:latest .

buildcn:
	docker build --build-arg GOPROXY=https://goproxy.cn --build-arg APKREPOSITORY=mirrors.tuna.tsinghua.edu.cn -t ${NAMESPACE}/${NAME}:latest .
	# docker build --progress=plain --build-arg GOPROXY=https://goproxy.cn --build-arg APKREPOSITORY=mirrors.tuna.tsinghua.edu.cn -t ${NAMESPACE}/${NAME}:latest .

full: build
	docker tag -t ${NAMESPACE}/${NAME}:${TAG} ${NAMESPACE}/${NAME}:latest

run:
	docker run --rm -it -p 8080:8080 ${NAMESPACE}/${NAME}
