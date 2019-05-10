#!/bin/bash
# start fabric containers
docker-compose -f deployment/docker-compose-peer1.yml up -d

# start cli container
docker-compose -f deployment/docker-compose-cli1.yml up -d