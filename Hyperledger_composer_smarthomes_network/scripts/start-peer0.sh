#!/bin/bash

# deploy peer0
docker-compose -f ../deployment/docker-compose-peer0.yml up -d 

# deploy cli0
docker-compose -f ../deployment/docker-compose-cli0.yml up -d 

