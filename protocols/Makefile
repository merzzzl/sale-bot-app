PROTO_REST_DIR = $(shell go list -m -f '{{.Dir}}' github.com/merzzzl/proto-rest-api)
OPENAPI_GENERATOR_VERSION = 7.25.0
DOCKER_USER ?= $(shell id -u):$(shell id -g)

.PHONY: generate generate-go generate-ts generate-doc

generate-go:
	go mod download
	mkdir -p go
	protoc -I=. -I "$(PROTO_REST_DIR)" --go_out=go --go-rest_out=go protobuff/*.proto

generate-ts:
	mkdir -p ts/src
	docker run --rm --user "$(DOCKER_USER)" -v "$(CURDIR):/local" \
		openapitools/openapi-generator-cli:v$(OPENAPI_GENERATOR_VERSION) generate \
		-i /local/go/api/v1_swagger.json -g typescript-axios -o /local/ts/src \
		--skip-validate-spec \
		--additional-properties=typescriptThreePlus=true,withInterfaces=true,modelPropertyNaming=camelCase
	npm ci
	npm run build

generate-doc:
	npx --yes swagger-markdown -i ./go/api/v1_swagger.json -o README.md.tmp
	sed '1,4d' README.md.tmp > README.md
	rm README.md.tmp

generate: generate-go generate-ts generate-doc
