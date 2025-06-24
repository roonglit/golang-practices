Install sqlc

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest


goose -dir=db/migrate -allow-missing postgres "host=localhost port=5432 user=postgres password=postgres dbname=learning sslmode=disable" up