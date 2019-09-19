#!/usr/bin/env bash

docker rm -f testalpine2
docker rm -f testalpine3

docker-compose -f docker-compose.yml up -d

docker logs -f testalpine2