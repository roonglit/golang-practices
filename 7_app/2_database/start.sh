#!/bin/sh

set -e

echo "run db migration"
goose --dir=db/migrate -allow-missing postgres $DB_URI up

echo "start server"
exec "$@"