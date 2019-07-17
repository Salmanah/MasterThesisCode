#!/bin/bash
# join peer0 to channel
# execute this command from server1

export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

docker exec cli peer channel fetch 0 mychannel.block -c mychannel -o orderer0.example.com:7050 --tls --cafile $ORDERER_CA

