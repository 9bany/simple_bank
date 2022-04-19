#!/usr/bin/env bash
set -e
echo "run db migrate"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "start server"
exec "$@"