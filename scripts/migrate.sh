#!/bin/bash

go install github.com/pressly/goose/v3/cmd/goose@latest

# Run goose to MySQL using envvars which are DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME
mysql -u $DB_USER -p$DB_PASSWORD -h $DB_HOST -e "CREATE DATABASE IF NOT EXISTS $DB_NAME"
goose -dir db/migrations mysql "$DB_USER:$DB_PASSWORD@($DB_HOST:$DB_PORT)/$DB_NAME" up
