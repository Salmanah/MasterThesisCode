#!/bin/bash

#Start peer0 container
docker-compose -f ../../deployment/docker-compose-peer3.yml up -d

#Start cli container 
docker-compose -f ../../deployment/docker-compose-cli3.yml up -d
