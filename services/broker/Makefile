build:
	@cd ./cmd/api && go build -o ../../bin/broker

build-push:
	docker build -f broker.dockerfile -t goocarry/broker:latest .

run: build
	./bin/broker