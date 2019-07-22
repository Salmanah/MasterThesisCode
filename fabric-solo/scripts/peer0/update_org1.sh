#!/bin/bash

##Used to update peerAnchor in org1


docker exec cli peer channel update -o orderer0.example.com:7050 -c mychannel -f network-config/Org1MSPanchors.tx  
 
