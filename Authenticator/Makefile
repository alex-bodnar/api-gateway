.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: all test build run run-build clean rebuild

APP_NAME := authenticator-microservice
VERSION := 0.0.1
TAG := $(shell git describe --abbrev=0 --tags)
COMMIT_DATE := $(shell git log -1 --date=format:"%y-%m-%dT%TZ" --format="%ad")
COMMIT_LAST := $(shell git rev-parse HEAD)
FORTUNE_COOKIE := $(shell curl -s http://yerkee.com/api/fortune/cookie  | \
					sed -e 's/\\t//g;s/\\n//g;s/[\#:%*@/{}\"\\]//g;s/fortune//g;s/\x27/ /g')

LDFLAGS=" \
       		-X 'main.appName=$(APP_NAME)' \
           	-X 'main.version=$(VERSION)' \
           	-X 'main.tag=$(TAG)' \
           	-X 'main.fortuneCookie=$(FORTUNE_COOKIE)' \
           	-X 'main.date=$(COMMIT_DATE)' \
           	-X 'main.commit=$(COMMIT_LAST)'"

all: run

test:
	go clean -testcache ./internal/...
	go test ./internal/...

build:
	cd cmd && go build -ldflags=$(LDFLAGS) -o ../build/$(APP_NAME) ./.

run:
	cd cmd; go run -ldflags=$(LDFLAGS) -race main.go -c ../volumes/config.yaml

run-build: mod build
	cd build/ && ./$(APP_NAME) -c ../volumes/config.yaml

clean:
	go clean ./...
	rm -r vendor build

rebuild: clean build

mod:
	go mod tidy

build-docker:
	go mod vendor;
	docker build \
	--build-arg APP_NAME=$(APP_NAME) \
	--build-arg VERSION=$(VERSION) \
	--build-arg COMMIT_DATE=$(COMMIT_DATE) \
	--build-arg FORTUNE_COOKIE="Test" \
	--build-arg TAG=$(TAG) \
	--build-arg COMMIT_LAST=$(COMMIT_LAST) \
	-t $(APP_NAME) .;
	rm -r vendor

swagger-generate:
	redoc-cli build docs/swagger.yaml -o docs/index.html

swagger-serve:
	redoc-cli serve docs/swagger.yaml
