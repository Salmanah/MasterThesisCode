#!/bin/bash

#Instantiate chaincode
export COLLECTIONS_PATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/network-config/

docker exec -it cli peer chaincode instantiate -o orderer0.example.com:7050  -C mychannel -c '{"Args":["Init"]}' -n device -v v0 -P "OR('Org1MSP.member','Org2.member')" --collections-config $COLLECTIONS_PATH/collections_config.json

