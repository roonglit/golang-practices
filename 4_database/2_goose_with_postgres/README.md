Run the postgresql using docker compose
docker compose up

Create a database on postgresql
PGPASSWORD=postgres createdb -h localhost -U postgres learning

Run Goose Command
goose postgres "host=localhost port=5432 user=postgres password=postgres dbname=learning sslmode=disable" up