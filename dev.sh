#!/bin/bash

# setup some env vards
export IS_PRODUCTION=false
export VERBOSE_LOGGING=true
export SERVER_PORT=8081
export HEALTH_NAME=turelo-api
export AWS_REGION=us-west-2
export NEGRONI_LOGGER_NAME=web
export OIDC_URL=https://apac-syd-partner01-test.apigee.net/.well-known/openid-configuration
export DATABASE_URL=postgres://postgres:postgres@127.0.0.1:5432/turelogo?sslmode=disable
export DATABASE_ENDPOINT=http://localhost:8000
export DATABASE_FEEDS_TABLE=turelo_dev_feeds_table
export DATABASE_FEEDENTRIES_TABLE=turelo_dev_feed_entries_table
export DATABASE_CATEGORIES_TABLE=turelo_dev_categories_table


# run the app
go run server.go

