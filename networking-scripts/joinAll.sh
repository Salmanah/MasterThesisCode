#!/bin/bash

echo "All peers join channel"
#Open 4 terminal tabs
tmux send -t CA.1  ./joinChannel_peer4.sh ENTER ##Peer4
tmux send -t CA.3  ./joinChannel_peer2.sh ENTER ##Peer2
tmux send -t CA.4  ./joinChannel_peer1.sh ENTER ##Peer1
tmux send -t CA.5  ./joinChannel_peer3.sh ENTER ##Peer3