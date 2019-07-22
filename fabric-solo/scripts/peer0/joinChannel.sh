#!/bin/bash
# join peer0 to channel
# execute this command from server1
docker exec cli peer channel join -b mychannel.block

# copy mychannel.block from peer0 to host(server1)
#docker cp peer0.org1.example.com:/mychannel.block .
