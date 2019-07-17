#!/bin/bash


echo "Opening 4 terminals for simulation"
#Open 4 terminal tabs
tmux new-session -s "CA" \; \
	send-keys 'ssh  ca' C-m \; \
	split-window -h -p 50 \; \
	send-keys 'ssh  peer0' C-m \; \
	split-window -v -p 50 \; \
	send-keys 'ssh peer1' C-m \; \
	select-pane -t 0 \; \
	split-window -v -p 50 \; \
	send-keys 'ssh peer2' C-m \; \
