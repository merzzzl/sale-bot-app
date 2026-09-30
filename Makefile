.PHONY: install build test lint docker
install:
	npm --prefix protocols ci
	npm --prefix protocols run build
	npm --prefix frontend ci
	cd backend && go mod download

build:
	npm --prefix protocols run build
	npm --prefix frontend run build
	cd backend && go build -o sale-bot-app ./cmd/app

test:
	cd protocols && go test ./...
	cd backend && go test ./...

lint:
	npm --prefix frontend run lint
	cd protocols && go vet ./...
	cd backend && go vet ./...

docker:
	docker build -t ghcr.io/merzzzl/sale-bot-app:latest .
