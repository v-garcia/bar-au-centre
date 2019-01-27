#!/bin/sh

set -e

until psql "${DATABASE_URL}" -q -c '\q'; do
    echo "Postgres is unavailable - sleeping"
    sleep 4s
done

echo "Postgres is up"

exec $exit