#!/bin/bash

echo "Generating crypto-material"
#Generate crypto-config
cryptogen generate --config=./crypto-config.yaml
mv crypto-config/ ..
../bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ../network-config/genesis.block --channelID syschannel
../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ../network-config/channel.tx -channelID mychannel
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ../network-config/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ../network-config/Org2MSPanchors.tx -channelID mychannel -asOrg Org2MSP
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ../network-config/Org3MSPanchors.tx -channelID mychannel -asOrg Org3MSP

