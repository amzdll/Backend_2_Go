include .env

LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	# Migration
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

	# Swagger
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest

migration-status:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} status -v

migration-up:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} up -v

migration-down:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} down -v

generate-swagger:
	${LOCAL_BIN}/swag init -g ${SWAGGER_SRC_DIR} -o ${SWAGGER_OUTPUT_DIR}

