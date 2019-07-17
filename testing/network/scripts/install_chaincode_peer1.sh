#!/bin/bash

#install chaincode on peer, cli container connects to peer
docker exec -e "CORE_PEER_ADDRESS=peer0.org2.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_TLS_ROOTCERT=/var/hyperledger/msp/peer/tls/ca.crt" -it cli peer chaincode install -n device -p github.com/chaincode -v v0
