PKG_DIR=pkg

#builds all services
build:
	go build -o ./bin/dndhub ./cmd/main.go

run: build
	./bin/dndhub

test:
	go test -v ./...

build-collector:
	@cd	$(PKG_DIR)/collector && echo "asd" && $(MAKE) build

run-collector: build-collector
	@cd $(PKG_DIR)/collector && ./bin/collector
