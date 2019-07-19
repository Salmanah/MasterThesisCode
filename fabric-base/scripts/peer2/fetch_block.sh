#!/bin/bash
# join peer0 to channel
# execute this command from server1


docker exec cli peer channel fetch 0 mychannel.block -c mychannel -o orderer0.example.com:7050 
