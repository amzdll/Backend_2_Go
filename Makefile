include .env

LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	# Migration
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

migration-status:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} status -v

migration-up:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} up -v

migration-down:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} down -v

generate-swagger:
	swag init -g src/cmd/shop/main.go -o src/api/shop/

