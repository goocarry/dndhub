build:
	@cd ./cmd/api && go build -o ../../bin/logger

build-push:
	docker build -f logger.dockerfile -t goocarry/logger:latest .

run: build
	./bin/logger