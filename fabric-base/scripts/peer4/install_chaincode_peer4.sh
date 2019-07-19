
#!/bin/bash

#install chaincode on peer, cli container connects to peer
docker exec -it cli peer chaincode install -n device -p github.com/chaincode -v v0
