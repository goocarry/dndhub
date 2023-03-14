include .env

migrate-up:
	@echo "asdasd"	
	@echo ${POSTGRESQL_URL} && cd ./db/migrations && migrate -database ${POSTGRESQL_URL} -path . up

build:
	@cd ./cmd/api && go build -o ./bin/auth