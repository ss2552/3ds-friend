SERVER_BUILD := friend

DATE_TIME    := $(shell date --iso=seconds)
BUILD_STRING := $(SERVER_BUILD), $(DATE_TIME)

SECURE_SERVER_HOST = 0.0.0.0
SECURE_SERVER_PORT = 60001
AUTHENTICATION_SERVER_PORT = 60000
POSTGRES_URI = postgres://username:password@localhost/friend?sslmode=disable
POSTGRES_MAX_CONNECTIONS = 1
HEALTH_CHECK_PORT = 0
AES_KEY = $(SERVER_BUILD)

.PHONY: build

build:
	go get -u
	go mod tidy
	go build -ldflags "-X 'main.serverBuildString=$(BUILD_STRING)'" -o ./build/$(SERVER_BUILD)

