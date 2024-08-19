#!/bin/bash

# Prompt the user for their GitHub token
read -p "Enter your DockerHub username: " DOCKERHUB_USERNAME
read -p "Enter your DockerHub password: " DOCKERHUB_PASSWORD
read -p "Enter the new version tag (e.g.: v1.0.1): " TAG

# Remove previous package builds
rm -rf cluster/my-cluster-*.xpkg

# Build the package
crossplane xpkg build -f ./cluster

# Login to DockerHub
docker login -u "$DOCKERHUB_USERNAME" -p "$DOCKERHUB_PASSWORD"

# Push the package to DockerHub
crossplane xpkg push -f ./cluster/my-cluster-*.xpkg index.docker.io/$DOCKERHUB_USERNAME/my-cluster:$TAG