#!/bin/bash

echo "Killing all docker containers and images"
#Open 4 terminal tabs
	tmux send -t CA.0  "docker system prune -a" ENTER ##Orderer
	tmux send -t CA.1  "docker system prune -a" ENTER ##Peer4
	tmux send -t CA.2  "docker system prune -a" ENTER ##Peer0
	tmux send -t CA.3  "docker system prune -a" ENTER ##Peer2
	tmux send -t CA.4  "docker system prune -a" ENTER #Peer1
	tmux send -t CA.5  "docker system prune -a" ENTER #Peer32

	



