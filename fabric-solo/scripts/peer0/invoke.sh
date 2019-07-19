#!/bin/bash


#Test invoke for peer
docker exec -it cli peer chaincode invoke -o orderer0.example.com:7050  -C mychannel -n device -c '{"Args":["sendDeviceReadingPrivate","BLACK","FRIDGE","DATA"]}'
