#!/bin/bash

#Instantiate chaincode
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

docker exec -it cli peer chaincode instantiate -o orderer0.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -n devicep github.com/chaincode -v v0 -c '{"Args":["initEnvironmentPrivate"]}' -P "OR('Org1MSP.member','Org2MSP.member')" --collections-config ../collections_config.json

