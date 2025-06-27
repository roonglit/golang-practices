DB_NAME ?= learning

db_create_dev:
	PGPASSWORD=postgres createdb -h localhost -U postgres $(DB_NAME)_development

db_create_test:
	PGPASSWORD=postgres createdb -h localhost -U postgres $(DB_NAME)_test

db_create:
	make db_create_dev
	make db_create_test

db_migrate_dev:
	goose -dir=db/migrate -allow-missing postgres "host=localhost port=5432 user=postgres password=postgres dbname=$(DB_NAME)_development sslmode=disable" up

db_migrate_test:
	goose -dir=db/migrate -allow-missing postgres "host=localhost port=5432 user=postgres password=postgres dbname=$(DB_NAME)_test sslmode=disable" up

db_migrate:
	make db_migrate_dev
	make db_migrate_test

g_migration_%:
	goose -dir=db/migrate create $* sql

sqlc:
	sqlc generate

test:
	GIN_MODE=test ginkgo ./...

server:
	go run main.go

PHONY: db_create db_migrate g_migration sqlc server