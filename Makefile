SERVICES_DIR=services

up:
	@echo "Starting docker images..."
	docker-compose up
	@echo "Docker images started!"

up-build:
	@echo "Stopping docker images (if running)..."
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

build:
	@echo "Stopping docker images (if running)..."
	docker-compose down
	@echo "Building docker images..."
	docker-compose build
	@echo "Docker images built!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

test:
	go test -v ./...

build-collector:
	@echo "Starting build collector service..."
	@cd	$(SERVICES_DIR)/collector && $(MAKE) build
	@echo "Done!"

run-collector: build-collector
	@cd $(SERVICES_DIR)/collector && ./bin/collector

build-broker:
	@echo "Starting build broker service..."
	@cd	$(SERVICES_DIR)/broker && $(MAKE) build
	@echo "Done!"

run-broker: build-broker
	@echo "Starting broker service..."
	@cd $(SERVICES_DIR)/broker && ./bin/broker
	@echo "Broker service is running..."

build-auth:
	@echo "Starting build broker service..."
	@cd $(SERVICES_DIR)/auth  && $(MAKE) build
	@echo "Done!"

run-auth:
	@echo "Starting auth service..."
	@cd $(SERVICES_DIR)/auth && ./bin/auth
	@echo "Auth service is running..."

migrate-db-auth:
	@echo "Starting migrate db auth service..."
	@cd $(SERVICES_DIR)/auth && $(MAKE) migrate-up
