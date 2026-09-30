commit:
	sh ./scripts/commit.sh

generate-go:
	rm -rf go
	mkdir go
	protoc \
		-I=. \
		-I ${GOPATH}/pkg/mod/github.com/merzzzl/proto-rest-api@v0.0.1-alpha.28 \
		--go_out=go \
		--go-rest_out=go \
		protobuff/*.proto

generate-ts:
	rm -rf ts
	mkdir ts
	npx @openapitools/openapi-generator-cli generate \
		-i go/api/v1_swagger.json \
		-g typescript-axios \
		-o ts/src \
		--skip-validate-spec \
		--additional-properties=typescriptThreePlus=true,withInterfaces=true,modelPropertyNaming=camelCase
	npm i
	npm run build

generate-doc:
	rm README.md
	npx swagger-markdown \
		-i ./go/api/v1_swagger.json \
		-o README.md.tmp
	sed '1,4d' README.md.tmp > README.md && rm README.md.tmp

generate: generate-go generate-ts generate-doc
