#!/usr/bin/env -S bash -e

echo "CREATE DATABASE \"shorten-db\" \
        OWNER = postgres \
        TEMPLATE=template0 \
        LC_COLLATE = 'C' \
        LC_CTYPE = 'C';" | psql 'host=127.0.0.1 sslmode=disable user=postgres'
