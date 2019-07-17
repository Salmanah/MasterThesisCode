#!/bin/bash
cryptogen generate --config=./crypto-config.yaml

sleep 3

mv crypto-config ../

sleep 3

../bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ../network-config/genesis.block -channelID syschannel

sleep 3

../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ..//network-config/channel.tx -channelID mychannel

sleep 3

../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ../network-config/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP

sleep 3

../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ../network-config/Org2MSPanchors.tx -channelID mychannel -asOrg Org2MSP

sleep 3

../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ../network-config/Org3MSPanchors.tx -channelID mychannel -asOrg Org3MSP