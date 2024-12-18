.PHONY: up down build migrate test help

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

up: ### Run docker-compose
	docker-compose up --build -d && docker-compose logs -f

down: ### Down docker-compose
	docker-compose down --remove-orphans

build:
	docker compose build

migrate:
	docker compose run backend go run db/migrate.go

test:
	docker compose run web_app go test ./...

# backend
generate:
	make generate-solution-api
	make generate-runner-api

generate-solution-api:
	mkdir -p backend/pkg/solution_v1
	protoc --proto_path proto --proto_path backend/vendor.protogen \
	--go_out=backend/pkg/solution_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=backend/bin/protoc-gen-go \
	--go-grpc_out=backend/pkg/solution_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=backend/bin/protoc-gen-go-grpc \
	--grpc-gateway_out=backend/pkg/solution_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=backend/bin/protoc-gen-grpc-gateway \
	--validate_out lang=go:backend/pkg/solution_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=backend/bin/protoc-gen-validate \
	--openapiv2_out=allow_merge=true,merge_file_name=api:backend/pkg/swagger \
	--plugin=protoc-gen-openapiv2=backend/bin/protoc-gen-openapiv2 \
	proto/backend.proto

generate-webapp:
	mkdir -p web_app/pkg/solution_v1
	protoc \
        --proto_path proto --proto_path backend/vendor.protogen \
        --go_out=web_app/pkg/solution_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=backend/bin/protoc-gen-go \
		--go-grpc_out=web_app/pkg/solution_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=backend/bin/protoc-gen-go-grpc \
		--grpc-gateway_out=web_app/pkg/solution_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=backend/bin/protoc-gen-grpc-gateway \
		--validate_out lang=go:web_app/pkg/solution_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=backend/bin/protoc-gen-validate \
		proto/backend.proto




generate-runner-api:
	mkdir -p backend/pkg/runner_v1
	protoc --proto_path proto --proto_path backend/vendor.protogen \
	--go_out=backend/pkg/runner_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=backend/bin/protoc-gen-go \
	--go-grpc_out=backend/pkg/runner_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=backend/bin/protoc-gen-go-grpc \
	--experimental_allow_unused_imports \
	proto/runner.proto

generate-templ:
	templ generate -path web_app

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/a-h/templ/cmd/templ@latest
