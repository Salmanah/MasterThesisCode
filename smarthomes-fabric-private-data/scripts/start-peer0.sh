#!/bin/bash

#Start peer0 container
docker-compose -f ../deployment/docker-compose-peer0.yml up -d

#Start cli container 
docker-compose -f ../deployment/docker-compose-cli0.yml up -d