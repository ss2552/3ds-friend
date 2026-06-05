RED    := $(shell tput setaf 1)
BLUE   := $(shell tput setaf 4)
CYAN   := $(shell tput setaf 14)
ORANGE := $(shell tput setaf 202)
YELLOW := $(shell tput setaf 214)
RESET  := $(shell tput sgr0)

ifeq ($(shell which go),)
# TODO - Read contents from .git folder instead?
$(error "$(RED)go command not found. Install go to continue $(BLUE)https://go.dev/doc/install$(RESET)")
endif

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

default:
	go get -u
	go mod tidy
	go build -ldflags "-X 'main.serverBuildString=$(BUILD_STRING)'" -o ./build/$(SERVER_BUILD)

