#!/bin/bash

#Instantiate chaincode
docker exec -it cli peer chaincode instantiate -o orderer0.example.com:7050  -C mychannel -c '{"Args":["Init"]}' -n device -v v0

