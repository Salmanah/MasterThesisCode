#!/bin/bash

echo "install chaincode on all peers"
#Open 4 terminal tabs
	tmux send -t CA.1  ./install_chaincode_peer4.sh ENTER ##Peer4
	tmux send -t CA.2  ./install_chaincode.sh ENTER ##Peer0
	tmux send -t CA.3  ./install_chaincode_peer2.sh ENTER ##Peer2
	tmux send -t CA.4  ./install_chaincode_peer1.sh ENTER #Peer1
	tmux send -t CA.5  ./install_chaincode_peer3.sh ENTER #Peer3

	



