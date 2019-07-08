#!/bin/bash
docker exec peer0.org3.example.com peer channel fetch 0 mychannel.block -c mychannel -o orderer0.example.com:7050

docker exec -e "CORE_PEER_LOCALMSPID=Org3MSP" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/users/Admin@org3.example.com/msp" peer0.org3.example.com peer channel join -b mychannel.block

