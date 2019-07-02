#!/bin/bash
# start fabric containers
docker-compose -f ../deployment/docker-compose-peer4.yml up -d

# start cli container
docker-compose -f ../deployment/docker-compose-cli4.yml up -d