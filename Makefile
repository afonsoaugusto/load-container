#!make

SHELL		       = bash

#git submodule add git@github.com:afonsoaugusto/base-ci.git
BASE_MAKEFILE := $(shell git submodule update --init --recursive)
include base-ci/Makefile

build: docker-build
scan: docker-scan
publish: docker-publish
