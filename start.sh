#!/bin/sh

set -e # script will exit immediately if command return non-zero

echo "start app"
exec "$@" # takes all parameters passed to the script and run it