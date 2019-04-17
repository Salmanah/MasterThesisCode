
# start fabric containers
docker-compose -f deployment/docker-compose-peer0.yml up -d

# start cli container
docker-compose -f deployment/docker-compose-cli0.yml up -d