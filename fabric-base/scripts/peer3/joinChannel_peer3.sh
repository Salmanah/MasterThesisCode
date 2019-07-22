#!/bin/bash

#docker cp mychannel.block peer0.org2.example.com:/mychannel.block
docker exec cli peer channel fetch 0 mychannel.block -c mychannel -o orderer0.example.com:7050 

docker exec cli peer channel join -b mychannel.block
