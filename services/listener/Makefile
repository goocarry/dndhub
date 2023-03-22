build:
	go build -o ./bin/listener

build-push:
	docker build -f listener.dockerfile -t goocarry/listener:latest .

run: build
	./bin/listener