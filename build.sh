#!/bin/bash

# Build and push the image to GitHub Container Registry

docker buildx build --platform=linux/amd64 -t ghcr.io/altepizza/mqtt_broadcaster:latest .
docker push ghcr.io/altepizza/mqtt_broadcaster:latest
