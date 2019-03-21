#!/bin/bash

# deploy peer0
docker-compose -f ../deployment/docker-compose-peer2.yml up -d 

# deploy cli0
docker-compose -f ../deployment/docker-compose-cli2.yml up -d 

