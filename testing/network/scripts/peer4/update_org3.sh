#!/bin/bash

##Used to update peerAnchor in org3

export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

docker exec -e "CORE_PEER_ADDRESS=peer0.org3.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org3MSP" -e "CORE_PEER_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp" cli peer channel update -o orderer0.example.com:7050 -c mychannel -f network-config/Org3MSPanchors.tx --tls --cafile $ORDERER_CA 
 
