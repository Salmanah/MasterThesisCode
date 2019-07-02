#!/bin/bash
docker exec peer1.org2.example.com peer channel fetch 0 mychannel.block -c mychannel -o orderer0.example.com:7050

docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/users/Admin@org2.example.com/msp" peer1.org2.example.com peer channel join -b mychannel.block

