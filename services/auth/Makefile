include .env

migrate-up:
	@echo ${POSTGRESQL_URL} && cd ./db/migrations && migrate -database ${POSTGRESQL_URL} -path . up

build:
	@cd ./cmd/api && go build -o ./bin/auth

build-push:
	docker build -f auth.dockerfile -t goocarry/auth:latest .
