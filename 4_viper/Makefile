db_create:
	mysql -u root -ppassword -h 127.0.0.1 -e "create database golang_development"

db_drop:
	mysql -u root -ppassword -h 127.0.0.1 -e "drop database golang_development"

generate_migration:
	migrate create -ext sql -dir db/migrate $(NAME)

db_migrate:
	migrate -path db/migrate -database "mysql://root:password@tcp(127.0.0.1:3306)/golang_development" -verbose up

db_migrate_down:
	migrate -path db/migrate -database "mysql://root:password@tcp(127.0.0.1:3306)/golang_development" -verbose down

sqlc:
	sqlc generate

credentials_edit:
	config_editor

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: db_create db_drop generate_migration db_migrate db_migrate_down sqlc test server