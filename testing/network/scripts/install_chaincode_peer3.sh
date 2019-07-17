#!/bin/bash

#install chaincode on peer, cli container connects to peer
docker exec -e "CORE_PEER_ADDRESS=peer1.org2.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_TLS_ROOTCERT=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp" -it cli peer chaincode install -n device -p github.com/chaincode -v v0
