#!/bin/bash

echo "All peers join channel"
#Open 4 terminal tabs
tmux send -t CA.1  ./update_org3.sh ENTER ##Peer4
tmux send -t CA.2  ./update_org1.sh ENTER ##Peer0
tmux send -t CA.4  ./update_org2.sh ENTER ##Peer1
