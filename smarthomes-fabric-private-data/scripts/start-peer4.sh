#!/bin/bash

#Start peer0 container
docker-compose -f ../deployment/docker-compose-peer4.yml up -d

#Start cli container 
docker-compose -f ../deployment/docker-compose-cli4.yml up -d
