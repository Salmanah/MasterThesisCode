#!
# start fabric containers
docker-compose -f deployment/docker-compose-peer3.yml up -d

# start cli container
docker-compose -f deployment/docker-compose-cli3.yml up -d