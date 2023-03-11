SERVICES_DIR=services

test:
	go test -v ./...

build-collector:
	@cd	$(SERVICES_DIR)/collector && echo "asd" && $(MAKE) build

run-collector: build-collector
	@cd $(SERVICES_DIR)/collector && ./bin/collector
