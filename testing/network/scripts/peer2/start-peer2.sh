#!/bin/bash

#Start peer0 container
docker-compose -f ../../deployment/docker-compose-peer2.yml up -d

#Start cli container 
docker-compose -f ../../deployment/docker-compose-cli2.yml up -d