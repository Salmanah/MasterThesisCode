#!/bin/bash




docker exec -it cli peer chaincode query -o orderer0.example.com:7050  -C mychannel -n device -c '{"Args":["readDevice","DEVICE_001"]}'
