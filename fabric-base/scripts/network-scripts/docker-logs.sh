#!/bin/bash

echo "Docker logs on all peer containers"
#Open 4 terminal tabs
	tmux send -t CA.0  "docker logs orderer0.example.com" ENTER ##Orderer
	tmux send -t CA.1  "docker logs peer0.org3.example.com" ENTER ##Peer4
	tmux send -t CA.2  "docker logs peer0.org1.example.com" ENTER ##Peer0
	tmux send -t CA.3  "docker logs peer1.org1.example.com" ENTER ##Peer2
	tmux send -t CA.4  "docker logs peer0.org2.example.com" ENTER #Peer1
	tmux send -t CA.5  "docker logs peer1.org2.example.com" ENTER #Peer3

	



