## Hyperledger Caliper


## How to run with TLS

step 1) npm install
step 2) run npm install for fabric-ca-client version 1.4
step 3) run node script/main -c benchmark/smarthomes-private/config.yaml -n network/fabric-v1.3/2org2peercouchdb/fabric-go-tls.json

## How to run without TLS

Do steps 1 and 2 from "How to run TLS"
step 3 run node script with the right files for this network

## If it does not work, try changing the cryptopath hfc-cv and then test
