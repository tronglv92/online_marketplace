include .env
OS := $(shell uname)
GOPATH:=$(shell go env GOPATH)
TARGET := $(firstword $(MAKECMDGOALS))
OPTIONAL_PARAM := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
API_PROTO_FILES=$(shell find api -name *.proto)
EXT_PROTO_FILES=$(shell find external -name *.proto)

# Export environment variables based on OS
ifeq ($(OS),Darwin)
    export $(shell sed 's/=.*//' .env)
else
	ifeq ($(OS),Linux)
		export $(shell cat .env | xargs)
	else
		export
	endif
endif

.PHONY: run
# run app service
run:
	go run cmd/main/main.go