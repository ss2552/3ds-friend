SERVER_BUILD := friend

DATE_TIME    := $(shell date --iso=seconds)
BUILD_STRING := $(SERVER_BUILD), $(DATE_TIME)


.PHONY: build

build:
	go get -u
	go mod tidy
	go build -ldflags "-X 'main.serverBuildString=$(BUILD_STRING)'" -o ./build/$(SERVER_BUILD)
