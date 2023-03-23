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

build-push-collector:
	@echo "Starting build collector service and pushing to docker registry..."
	@cd	$(SERVICES_DIR)/collector && $(MAKE) build-push
	@echo "Done!"

run-collector: build-collector
	@cd $(SERVICES_DIR)/collector && ./bin/collector

build-broker:
	@echo "Starting build broker service..."
	@cd	$(SERVICES_DIR)/broker && $(MAKE) build
	@echo "Done!"

build-push-broker:
	@echo "Starting build broker service and pushing to docker registry..."
	@cd	$(SERVICES_DIR)/broker && $(MAKE) build-push
	@echo "Done!"

run-broker: build-broker
	@echo "Starting broker service..."
	@cd $(SERVICES_DIR)/broker && ./bin/broker
	@echo "Broker service is running..."

build-auth:
	@echo "Starting build auth service..."
	@cd $(SERVICES_DIR)/auth && $(MAKE) build
	@echo "Done!"

build-push-auth:
	@echo "Starting build auth service and pushing to docker registry..."
	@cd $(SERVICES_DIR)/auth && $(MAKE) build-push
	@echo "Done!"

run-auth:
	@echo "Starting auth service..."
	@cd $(SERVICES_DIR)/auth && ./bin/auth
	@echo "Auth service is running..."

migrate-db-auth:
	@echo "Starting migrate db auth service..."
	@cd $(SERVICES_DIR)/auth && $(MAKE) migrate-up

build-logger:
	@echo "Starting build logger service..."
	@cd	$(SERVICES_DIR)/logger && $(MAKE) build
	@echo "Done!"

build-push-logger:
	@echo "Starting build logger service and pushing to docker registry..."
	@cd	$(SERVICES_DIR)/logger && $(MAKE) build-push
	@echo "Done!"

run-logger: build-logger
	@echo "Starting logger service..."
	@cd $(SERVICES_DIR)/logger && ./bin/logger
	@echo "Logger service is running..."

build-mailer:
	@echo "Starting build mailer service..."
	@cd $(SERVICES_DIR)/mailer && $(MAKE) build
	@echo "Done!"

build-push-mailer:
	@echo "Starting build mailer service and pushing to docker registry..."
	@cd $(SERVICES_DIR)/mailer && $(MAKE) build-push
	@echo "Done!"

run-mailer:
	@echo "Starting mailer service..."
	@cd $(SERVICES_DIR)/mailer && ./bin/mailer
	@echo "Mailer service is running..."

build-listener:
	@echo "Starting build listener service..."
	@cd $(SERVICES_DIR)/listener && $(MAKE) build
	@echo "Done!"

build-push-listener:
	@echo "Starting build listener service and pushing to docker registry..."
	@cd $(SERVICES_DIR)/listener && $(MAKE) build-push
	@echo "Done!"

run-listener:
	@echo "Starting listener service..."
	@cd $(SERVICES_DIR)/listener && ./bin/listener
	@echo "Listener service is running..."

build-push-all: build-push-listener build-push-mailer build-push-auth build-push-logger build-push-broker
	@echo "All docker images are built and pushed to registry!"