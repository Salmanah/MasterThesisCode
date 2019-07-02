#!/bin/bash
# start fabric containers
docker-compose -f ../deployment/docker-compose-peer2.yml up -d

# start cli container
docker-compose -f ../deployment/docker-compose-cli2.yml up -d