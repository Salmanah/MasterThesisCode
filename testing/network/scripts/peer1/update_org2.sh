#!/bin/bash

##Used to update peerAnchor in org2

export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

docker exec cli peer channel update -o orderer0.example.com:7050 -c mychannel -f network-config/Org2MSPanchors.tx --tls --cafile $ORDERER_CA
