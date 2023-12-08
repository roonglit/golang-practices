* Run docker using `docker compose up`
* Calling mysql to create a database
** mysql -u root -h 127.0.0.1 -p
** using password as a password
** create database golang_dev;
* Install migrate
** go to CLI documentation https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
* Create a new folder db/migrate
* Calling migrate to create a new migration file
** migrate create -ext sql -dir db/migrate create_posts
* Create a makefile for creating a database
* Call migrate up
** migrate -path db/migrate -database "mysql://root:password@tcp(127.0.0.1:3306)/golang_development" -verbose up
** create posts & users migration
* Update Makefile
* Setting up sqlc
** brew install sqlc
** sqlc init