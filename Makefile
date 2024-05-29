NAME := gin-report
TAG := $(shell git log -1 --pretty=%h)
NAMESPACE := sidecus

.PHONY : build run

all : build run

build :
	docker build -t ${NAMESPACE}/${NAME}:latest .

run :
	docker run --rm -it -p 8080:8080 ${NAMESPACE}/${NAME}

fullbuild:
	docker build -t ${NAMESPACE}/${NAME}:${TAG} -t ${NAMESPACE}/${NAME}:latest .
