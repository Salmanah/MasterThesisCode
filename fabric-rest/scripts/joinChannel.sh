#!/bin/bash
# join peer0 to channel
# execute this command from server1
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b ../mychannel.block

# copy mychannel.block from peer0 to host(server1)
docker cp peer0.org1.example.com:/mychannel.block .
