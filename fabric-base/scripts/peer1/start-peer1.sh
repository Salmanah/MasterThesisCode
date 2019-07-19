#!/bin/bash

#Start peer0 container
docker-compose -f ../../deployment/docker-compose-peer1.yml up -d

#Start cli container 
docker-compose -f ../../deployment/docker-compose-cli1.yml up -d