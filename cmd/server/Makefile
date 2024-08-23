# Define variables
DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_FILE = ../../docker-compose.yml
SERVICE_API = api

# Build Docker images
.PHONY: build
build:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) build --no-cache --pull

# Start Docker containers
.PHONY: up
up:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d

# Stop Docker containers
.PHONY: down
down:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

# Restart Docker containers
.PHONY: restart
restart: down up

# Log into the Go container
.PHONY: shell
shell:
	docker exec -it $(SERVICE_API) bash

# Run API tests without Docker
.PHONY: test-local
test-local:
	go test ./src/tests/auth/login -v

# Run API tests inside the Docker container
.PHONY: test
test:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) exec $(SERVICE_API) go test ./src/tests/... -v

# Run tests from a specific test repository
.PHONY: test-repo
test-repo:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) exec $(SERVICE_API) go test ./src/tests/auth/login -v

# Run a single test file
.PHONY: test-file
test-file:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) exec $(SERVICE_API) go test ./src/tests/auth/login -v -run Test_can_login_if_valid_email_and_password_are_provided

# Show the status of Docker containers
.PHONY: ps
ps:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) ps

# Build and start the containers, then run tests
.PHONY: build-and-test
build-and-test: build up test

# Clean up Docker images and containers
.PHONY: clean
clean:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down --volumes --rmi all

# Display help message
.PHONY: help
help:
	@echo "Makefile for Dockerized Go API"
	@echo "Usage:"
	@echo "  make build           Build Docker images"
	@echo "  make up              Start Docker containers"
	@echo "  make down            Stop Docker containers"
	@echo "  make restart         Restart Docker containers"
	@echo "  make shell           Log into the Go container"
	@echo "  make test-local      Run API tests without Docker"
	@echo "  make test            Run API tests inside Docker container"
	@echo "  make test-repo       Run tests from a specific test repository"
	@echo "  make test-file       Run a single test file"
	@echo "  make ps              Show Docker container status"
	@echo "  make build-and-test  Build and start containers, then run tests"
	@echo "  make clean           Clean up Docker images and containers"
	@echo "  make help            Display this help message"
