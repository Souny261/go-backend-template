# Go parameters
GOCMD=go
GORUN=$(GOCMD) run

# Docker parameters
DOCKER_IMAGE=backend
DOCKER_CONTAINER=backend-container
DOCKER_PORT=3000

.PHONY: run build docker-run docker-stop docker-clean

# Run the application
run:
	$(GORUN) ./cmd/api

# Build Docker image
build:
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
run-docker:
	docker run -d --name $(DOCKER_CONTAINER) -p $(DOCKER_PORT):$(DOCKER_PORT) $(DOCKER_IMAGE)

# Stop Docker container
stop:
	docker stop $(DOCKER_CONTAINER)

# Remove Docker container
clean:
	docker rm $(DOCKER_CONTAINER)
