#!/bin/bash

#Instantiate chaincode
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export COLLECTIONS_PATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/network-config/

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" -it cli peer chaincode instantiate -o orderer0.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -c '{"Args":["initEnvironmentPrivate"]}' -n device -v 0.1 -P "OR('Org1MSP.member','Org2MSP.member','Org3MSP.member')" --collections-config $COLLECTIONS_PATH/collections_config.json

