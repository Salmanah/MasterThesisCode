#!/bin/bash

#install chaincode on peer, cli container connects to peer
docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_TLS_ROOTCERT=/var/hyperledger/msp/peer/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/msp/peer/msp/" -it cli peer chaincode install -n device -p github.com/chaincode -v v0
