#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -ev

# Shut down the Docker containers that might be currently running.
#stop all containers:**
docker kill $(docker ps -q)

#remove all containers**
docker rm $(docker ps -a -q)

