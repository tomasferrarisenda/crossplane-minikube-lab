#!/bin/bash

# Prompt the user for their GitHub token
read -p "Enter your AWS access key ID: " AWS_ACCESS_KEY_ID
read -p "Enter your AWS secret access key: " AWS_SECRET_ACCESS_KEY

crossplane xpkg build -f ./cluster

docker login -u "tferrari92" -p "password"

crossplane xpkg push -f ./cluster/aatt-cluster-*.xpkg index.docker.io/tferrari92/aatt-cluster:v1.0.0