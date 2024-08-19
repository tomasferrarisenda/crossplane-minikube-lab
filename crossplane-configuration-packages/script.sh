#!/bin/bash

# Prompt the user for their GitHub token
read -p "Enter your DockerHub username: " DOCKERHUB_USERNAME
read -p "Enter your DockerHub password: " DOCKERHUB_PASSWORD
read -p "Enter the new version tag (e.g.: v1.0.1): " TAG

rm -rf cluster/my-cluster-*.xpkg

crossplane xpkg build -f ./cluster

docker login -u "$DOCKERHUB_USERNAME" -p "$DOCKERHUB_PASSWORD"

crossplane xpkg push -f ./cluster/my-cluster-*.xpkg index.docker.io/$DOCKERHUB_USERNAME/my-cluster:$TAG