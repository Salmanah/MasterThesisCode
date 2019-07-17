#!/bin/bash

#Instantiate chaincode
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export COLLECTIONS_PATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/network-config/

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_TLS_ROOTCERT=/var/hyperledger/msp/peer/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/msp/peer/msp/" -it cli peer chaincode instantiate -o orderer0.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -c '{"Args":["Init"]}' -n device -v v0 -P "OR('Org1MSP.member','Org2.member')" --collections-config $COLLECTIONS_PATH/collections_config.json

