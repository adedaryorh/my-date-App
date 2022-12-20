SHELL:=bash
POSTGRES_DB:=postgres://postgres:root@127.0.0.1:5432/celebut_db?sslmode=disable
ifndef MIGRATION_STEPS
override MIGRATION_STEPS = 1
endif


# runtime options
COMMIT_HASH := $(shell git rev-parse --short HEAD)

LDFLAGS     := -X "main.commitHash=$(COMMIT_HASH)"

bash: ## Run bash inside app container
	@docker-compose exec app bash

.build-migrate:
	go build -mod=vendor -ldflags '$(LDFLAGS)' -tags migrate -o bin/celebut-migrate ./cmd/celebutapi/

build: clean .build-migrate ## Build binaries
	go build -mod=vendor -ldflags '$(LDFLAGS)' -o bin/celebut-api ./cmd/celebutapi/

build-static: ## Build binaries statically
	CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -mod=mod -v -a -installsuffix cgo -o bin/celebut-api ./cmd/celebutapi/

build-migrate-static:
	go build -mod=mod -ldflags '$(LDFLAGS)' -tags migrate -o bin/celebut-migrate ./cmd/celebutapi/

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

migrate-create:
	migrate create -ext sql -dir migrations $(name)
.PHONY: migrate-create

migrate-up:
	migrate -path migrations -database '$(POSTGRES_DB)' up
.PHONY: migrate-up

migrate-force:
	migrate -path migrations -database '$(POSTGRES_DB)' force $(migration)
.PHONY: migrate-up

migrate-down:
	migrate -path migrations -database '$(POSTGRES_DB)' down $(MIGRATION_STEPS)
.PHONY: migrate-down

run: ## Start the containers and attach it
	@docker-compose up -d

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

swag-fmt: ### swag format
	swag fmt
.PHONY: swag-fmt

stop: ## Stop any running container
	@docker-compose stop

test-unit: ## Execute unit tests
	@ginkgo --focus=${FOCUS} -cover --output-dir '.' -coverprofile 'cover.out' -race ./internal/... ./pkg/...
	find ./*/ -name "cover.out" -exec rm {} \;

tools: ## Install development tools
	go install github.com/onsi/ginkgo/v2/ginkgo@latest
	go install github.com/codegangsta/gin@latest
	go install github.com/swaggo/swag/cmd/swag@latest

tools.env: ## Copy .env.dist to .env if it does not exist yet
	@cp -n .env.dist .env 2> /dev/null || true

up: ## Start the containers and attach it
	@docker-compose up