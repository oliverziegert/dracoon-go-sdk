.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: swagger
swagger: ## Generate OpenAPI models
	$(MAKE) swagger-dracoon

.PHONY: swagger-dracoon
swagger-dracoon: ## Generate OpenAPI models for DRACOON
	openapi-generator generate -i https://dracoon.team/api/spec_v4 -g go -o ./client/dracoon/core/openapi
	find ./client/dracoon/core/openapi -mindepth 2 -delete
	find ./client/dracoon/core/openapi -not -name '*.go' -delete
	$(MAKE) fmt

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: ## Run tests.
	go test ./... -coverprofile cover.out
