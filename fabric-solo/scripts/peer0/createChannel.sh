#!/bin/bash

# create channel from peer0 on server1
# it connects to orderer0
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" cli peer channel create -o orderer0.example.com:7050 -c mychannel -f network-config/channel.tx 
