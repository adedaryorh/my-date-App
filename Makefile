SHELL:=bash

# runtime options
COMMIT_HASH := $(shell git rev-parse --short HEAD)

LDFLAGS     := -X "main.commitHash=$(COMMIT_HASH)"

bash: ## Run bash inside app container
	@docker-compose exec app bash

build: clean ## Build binaries
	go build -mod=vendor -ldflags '$(LDFLAGS)' -o bin/celebut-api ./cmd/celebutapi/

build-static: ## Build binaries statically
	CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -v -a -installsuffix cgo -o bin/celebut-api ./cmd/celebutapi/

clean: ## Cleanup runtime files
	rm -rf celebut-api *.coverprofile *.out

clean-all: clean ## Cleanup ALL runtime files
	rm -rf vendor

deps: ## Install dependencies
	go mod download

generate: ## Generate code using directives
	go generate ./...

init: tools.env
	@docker-compose build

run: ## Start the containers and attach it
	@docker-compose up -d

stop: ## Stop any running container
	@docker-compose stop

test-unit: ## Execute unit tests
	@ginkgo --focus=${FOCUS} -cover --output-dir '.' -coverprofile 'cover.out' -race ./internal/... ./pkg/...
	find ./*/ -name "cover.out" -exec rm {} \;

tools: ## Install development tools
	go install github.com/onsi/ginkgo/v2/ginkgo@latest
	go install github.com/codegangsta/gin@latest

tools.env: ## Copy .env.dist to .env if it does not exist yet
	@cp -n .env.dist .env 2> /dev/null || true

up: ## Start the containers and attach it
	@docker-compose up