MAKEFLAGS  	+= --silent
SHELL      	 = /bin/bash

PROJECT_NAME := ${CIRCLE_PROJECT_REPONAME}
BRANCH_NAME  := ${CIRCLE_BRANCH}
COMMIT_SHA   := $(shell echo ${CIRCLE_SHA1}| head -c 7)

ifndef CIRCLE_PROJECT_REPONAME
	PROJECT_NAME   			:= $(shell basename $(CURDIR))
	BRANCH_NAME   			:= $(shell git rev-parse --abbrev-ref HEAD)
	COMMIT_SHA   			:= $(shell git rev-parse --short HEAD)
endif

# IMAGE_REGISTRY					 := 
# IMAGE_REGISTRY_USERNAME  := 
# IMAGE_REGISTRY_TOKEN		 := 

IMAGE_NAME_COMMIT  := ${IMAGE_REGISTRY_USERNAME}/${PROJECT_NAME}:${COMMIT_SHA}-${BRANCH_NAME}
IMAGE_NAME_BRANCH  := ${IMAGE_REGISTRY_USERNAME}/${PROJECT_NAME}:${BRANCH_NAME}
IMAGE_NAME_LATEST  := ${IMAGE_REGISTRY_USERNAME}/${PROJECT_NAME}:latest

docker-login:
	echo ${IMAGE_REGISTRY_TOKEN} | docker login -u ${IMAGE_REGISTRY_USERNAME} --password-stdin ${IMAGE_REGISTRY}

docker-build: docker-login
	docker build -t $(IMAGE_NAME_COMMIT) . && \
	docker push $(IMAGE_NAME_COMMIT) && \
	docker logout

docker-publish-branch: docker-login
	docker pull $(IMAGE_NAME_COMMIT)  && \
	docker tag $(IMAGE_NAME_COMMIT) $(IMAGE_NAME_BRANCH)  && \
	docker push $(IMAGE_NAME_BRANCH)

docker-publish: docker-login
	docker pull $(IMAGE_NAME_COMMIT)  && \
	docker tag $(IMAGE_NAME_COMMIT) $(IMAGE_NAME_LATEST)  && \
	docker push $(IMAGE_NAME_LATEST) && \
	docker logout

docker-scan:
	trivy --exit-code 1 $(IMAGE_NAME_COMMIT)

docker-clear:
	docker logout

docker-branch: docker-build docker-publish-branch docker-clear

build: docker-build
scan: docker-scan
publish: docker-publish
