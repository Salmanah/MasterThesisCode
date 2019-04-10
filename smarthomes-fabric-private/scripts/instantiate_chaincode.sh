#!/bin/bash

#Instantiate chaincode 
docker exec -it cli peer chaincode instantiate -o orderer0.example.com:7050 -C mychannel -n device github.com/chaincode -v v0 -c '{"Args":["initDevice","DEVICE_001","100","HOME_001","SALMAN"]}'

