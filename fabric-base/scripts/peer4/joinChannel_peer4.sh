#!/bin/bash

#docker cp mychannel.block peer0.org2.example.com:/mychannel.block

docker exec cli peer channel join -b mychannel.block
