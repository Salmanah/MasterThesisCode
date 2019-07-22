#!/bin/bash

#docker cp mychannel.block peer0.org2.example.com:/mychannel.block

export ORDERER_CA=/var/hyperledger/msp/orderer/msp/tlscacerts/tlsca.example.com-cert

docker exec cli peer channel fetch 0 mychannel.block -c mychannel -o orderer0.example.com:7050 --tls --cafile $ORDERER_CA
docker exec cli peer channel join -b mychannel.block
