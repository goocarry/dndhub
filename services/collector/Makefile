build:
	go build -o ./bin/collector

build-push:
	docker build -f collector.dockerfile -t goocarry/collector:latest .

run: build
	./bin/collector