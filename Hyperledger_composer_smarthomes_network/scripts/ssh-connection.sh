#!/bin/bash


echo "Opening 4 terminals for simulation"
#Open 4 terminal tabs
tmux new-session -s "CA" \; \
    send-keys 'ssh -i ca' C-m \; \
    split-window -h -p 50 \; \
    send-keys 'ssh -i  peer0' C-m \; \
    split-window -v -p 50 \; \
    send-keys 'ssh -i peer1' C-m \; \
    select-pane -t 0 \; \
    split-window -v -p 50 \; \
    send-keys 'ssh -i peer2' C-m \; \
