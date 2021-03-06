#!/bin/bash

echo "Opening 4 terminals for simulation"
#Open 4 terminal tabs
tmux new-session -s "CA" \; \
	send-keys 'ssh  ca' C-m \; \
	send-keys 'cd test/network/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-orderer.sh' C-m \; \
	split-window -v -p 50 \; \
	send-keys 'ssh  peer0' C-m \; \
	send-keys 'cd test/network/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-peer0.sh' C-m \; \
	split-window -h -p 50 \; \
	send-keys 'ssh peer1' C-m \; \
	send-keys 'cd test/network/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-peer1.sh' C-m \; \

	



