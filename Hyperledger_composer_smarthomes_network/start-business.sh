#!/bin/bash

#install business network

composer network install --card PeerAdmin@hlfv1 --archiveFile MasterThesisInPlayground/dist/smarthomes@0.0.17.bna

sleep 5

composer network start --networkName smarthomes --networkVersion 0.0.17 --networkAdmin admin --networkAdminEnrollSecret adminpw --card PeerAdmin@hlfv1 --file  networkadmin.card

sleep 5

composer card import --file networkadmin.card 




