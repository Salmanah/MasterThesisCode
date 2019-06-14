#!/bin/bash
docker cp mychannel.block peer0.org2.example.com:/mychannel.block

docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/users/Admin@org2.example.com/msp" peer0.org2.example.com peer channel join -b ../mychannel.block

rm mychannel.block