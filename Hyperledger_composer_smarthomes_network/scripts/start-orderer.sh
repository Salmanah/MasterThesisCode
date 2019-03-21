#!/bin/bash            

#Deploy ca, orderer and kafka
docker-compose -f ../deployment/docker-compose-kafka.yml up -d 