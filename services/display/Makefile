build:
	go build -o ./bin/display

build-push:
	docker build -f display.dockerfile -t goocarry/display:latest .

run: build
	./bin/display