Run aynqmon
```
docker run --rm --platform=linux/amd64 -p 8080:8080 \
  hibiken/asynqmon \
  --redis-addr=host.docker.internal:6379
```