build:
	@cd ./cmd/api && go build -o ../../bin/mailer

build-push:
	docker build -f mailer.dockerfile -t goocarry/mailer:latest .

run: build
	./bin/mailer