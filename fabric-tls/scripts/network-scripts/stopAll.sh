#!/bin/bash

echo "Stopping all docker containers"
#Open 4 terminal tabs
	tmux send -t CA.0  ./stop.sh ENTER
	tmux send -t CA.1  ./stop.sh ENTER
	tmux send -t CA.2  ./stop.sh ENTER
	tmux send -t CA.3  ./stop.sh ENTER
	tmux send -t CA.4  ./stop.sh ENTER
	tmux send -t CA.5  ./stop.sh ENTER
