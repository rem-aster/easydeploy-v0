.PHONY: up down build migrate test

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

migrate:
	docker compose run backend go run db/migrate.go

test:
	docker compose run web_app go test ./...

# backend
generate:
	make generate-solution-api

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

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
