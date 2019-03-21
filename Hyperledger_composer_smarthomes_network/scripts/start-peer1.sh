#!/bin/bash

# deploy peer0
docker-compose -f ../deployment/docker-compose-peer1.yml up -d 

# deploy cli0
docker-compose -f ../deployment/docker-compose-cli1.yml up -d 

