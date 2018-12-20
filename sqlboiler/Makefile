DB_HOST ?= localhost
DB_USER ?= postgres
DB_PWD ?= password
DB_PORT ?= 5432
DB_NAME ?= test

.PHONY: migrate build clean

build:

migrate:
	migrate -path ./sql -database 'postgres://$(DB_USER):$(DB_PWD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable' up

clean:
