include .env

LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	# Migration
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

get-deps:
	# Environment configuration
	go get github.com/joho/godotenv

	# Handler
	go get -u github.com/go-chi/chi/v5
	go get -u github.com/go-playground/validator/v10
	go get -u github.com/nicklaw5/go-respond

	# Database
	go get -u github.com/google/uuid
	go get -u github.com/jackc/pgx/v5
	go get -u github.com/georgysavva/scany/v2
	go get -u github.com/georgysavva/scany/v2/pgxscan

migration-status:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} status -v

migration-up:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} up -v

migration-down:
	${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN} down -v