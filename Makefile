build:
	go build -o ./bin/dndhub ./cmd/main.go

run: build
	./bin/dndhub

test:
	go test -v ./...