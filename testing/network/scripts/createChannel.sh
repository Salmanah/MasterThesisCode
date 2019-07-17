#!/bin/bash

# create channel from peer0 on server1
# it connects to orderer0
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" cli peer channel create -o orderer0.example.com:7050 -c mychannel -f network-config/channel.tx --tls --cafile $ORDERER_CA
