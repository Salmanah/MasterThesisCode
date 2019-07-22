#!/bin/bash

echo "Starting all docker containers"
#Open 4 terminal tabs
	tmux send -t CA.0  ./start-orderer.sh ENTER ##Orderer
	tmux send -t CA.1  ./start-peer4.sh ENTER ##Peer4
	tmux send -t CA.2  ./start-peer0.sh ENTER ##Peer0
	tmux send -t CA.3  ./start-peer2.sh ENTER ##Peer2
	tmux send -t CA.4  ./start-peer1.sh ENTER ##Peer1
	tmux send -t CA.5  ./start-peer3.sh ENTER ##Peer3