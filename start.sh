#!/bin/sh

set -e # script will exit immediately if command return non-zero

echo "run db migartion"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" --verbose up

echo "start app"
exec "$@" # takes all parameters passed to the script and run it