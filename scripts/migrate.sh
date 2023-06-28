#!/bin/bash

go install github.com/pressly/goose/v3/cmd/goose@latest

# Set default values for envvars used if not set
DB_USER=${DB_USER:-root}
DB_PASSWORD=${DB_PASSWORD:-password}
DB_HOST=${DB_HOST:-127.0.0.1}
DB_PORT=${DB_PORT:-3306}
DB_NAME=${DB_NAME:-diary}

# Run goose to MySQL using envvars which are DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME
mysql -u $DB_USER -p$DB_PASSWORD -h $DB_HOST -e "CREATE DATABASE IF NOT EXISTS $DB_NAME"
goose -dir db/migrations mysql "$DB_USER:$DB_PASSWORD@($DB_HOST:$DB_PORT)/$DB_NAME" up
