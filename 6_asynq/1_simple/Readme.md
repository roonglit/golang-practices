### Start Redis
- `docker run -d -p 6379:6379 redis:latest`
### Initialize Go module
- `go mod init asynq-demo`
- `go get github.com/hibiken/asynq`
### Start
- `go run main.go`