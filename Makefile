APP_NAME=example-be
DEFAULT_PORT=8100
.PHONY: setup init build dev test db-migrate-up db-migrate-down

setup:
	cd ~ && go get -v github.com/rubenv/sql-migrate/...
	cd ~ && go get github.com/golang/mock/gomock
	cd ~ && go get github.com/golang/mock/mockgen
	cp .env.sample .env && vim .env

build:
	env GOOS=darwin GOARCH=amd64 go build -o bin/server $(shell pwd)/cmd/server/

dev:
	go run ./cmd/server/main.go

admin:
	go run ./cmd/admin/*.go

test:
	go test -cover ./...

docker-build:
	docker build \
	--build-arg DEFAULT_PORT="${DEFAULT_PORT}" \
	-t ${APP_NAME}:latest .