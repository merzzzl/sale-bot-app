APP_NAME := $(notdir $(CURDIR))
IMAGE ?= ghcr.io/merzzzl/$(APP_NAME):local
GO_MODULES := backend $(if $(wildcard protocols/go.mod),protocols)
WEB_MODULES := $(if $(wildcard protocols/package.json),protocols) frontend

.PHONY: install install-go install-web check check-go check-web check-compose test lint build docker

install: install-go install-web

install-go:
	@set -eu; for module in $(GO_MODULES); do (cd $$module && go mod download); done

install-web:
	@set -eu; for module in $(WEB_MODULES); do npm --prefix $$module ci; if [ "$$module" = protocols ]; then npm --prefix $$module run build; fi; done

check: check-go check-web check-compose

check-go:
	@set -eu; for module in $(GO_MODULES); do (cd $$module && go vet ./... && go test -race ./... && go build ./...); done

check-web:
	npm --prefix frontend run lint
	npm --prefix frontend run typecheck
	npm --prefix frontend run build

check-compose:
	sh scripts/check-compose.sh
	@set -eu; for script in scripts/*.sh; do sh -n "$$script"; done
	python3 scripts/tests/test_bootstrap.py

test:
	@set -eu; for module in $(GO_MODULES); do (cd $$module && go test -race ./...); done

lint:
	@set -eu; for module in $(GO_MODULES); do (cd $$module && go vet ./...); done
	npm --prefix frontend run lint

build:
	@if [ -f protocols/package.json ]; then npm --prefix protocols run build; fi
	npm --prefix frontend run build
	cd backend && go build -trimpath -o app ./cmd/app

docker:
	docker build -t $(IMAGE) .
