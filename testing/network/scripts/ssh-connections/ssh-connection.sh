#!/bin/bash

echo "Opening 4 terminals for simulation"
#Open 4 terminal tabs
tmux new-session -s "CA" \; \
	send-keys 'ssh  ca' C-m \; \
	send-keys 'cd smarthomes-fabric-private-data/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-orderer.sh' C-m \; \
	split-window -v -p 75 \; \
	send-keys 'ssh  peer0' C-m \; \
	send-keys 'cd smarthomes-fabric-private-data/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-peer0.sh' C-m \; \
	split-window -h -p 50 \; \
	send-keys 'ssh peer1' C-m \; \
	send-keys 'cd smarthomes-fabric-private-data/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-peer1.sh' C-m \; \
	split-window -v -p 50 \; \
	send-keys 'ssh peer3' C-m \; \
	send-keys 'cd smarthomes-fabric-private-data/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-peer3.sh' C-m \; \
	select-pane -t 1 \; \
	split-window -v -p 50 \; \
	send-keys 'ssh peer2' C-m \; \
	send-keys 'cd smarthomes-fabric-private-data/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-peer2.sh' C-m \; \
	select-pane -t 0 \; \
	split-window -h -p 50 \; \
	send-keys 'ssh peer4' C-m \; \
	send-keys 'cd smarthomes-fabric-private-data/scripts' C-m \; \
	send-keys './stop.sh' C-m \; \
	send-keys './start-peer4.sh' C-m \; \

	



