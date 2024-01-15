#!/bin/sh

# Path: deployments/restart-service.sh

# Use this script to restart the service on the server.
# Place this script in the project root directory.
# Run this script from the project root directory.

# Pull the latest changes from the repository.
git pull

# build the project. Here go-gin-boilerplate is the name of the container.
docker exec go-gin-boilerplate make build

# Restart the service. Here go-gin-boilerplate is the name of the container.
docker exec go-gin-boilerplate ./build/ggb server
