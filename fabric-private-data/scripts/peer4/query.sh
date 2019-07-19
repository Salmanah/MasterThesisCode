#!/bin/bash


export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


docker exec -it cli peer chaincode query -o orderer0.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -n device -c '{"Args":["readDevice","DEVICE_001"]}'
