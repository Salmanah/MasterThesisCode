This is the readme file for the Business Network Definition created in Playground

# smarthomes

Blockchain environment for smart homes. 

# Change the directory to smarthomes dist subfolder
cd dist

# Create the archive
composer archive create  --sourceType dir --sourceName ../


# Install the BNA
composer network install -a ./smarthomes@0.0.1.bna -c PeerAdmin@hlfv1

# Start the BNA
composer network start -c PeerAdmin@hlfv1 -n smarthomes -V 0.0.1  -A admin -S adminpw

# Import the card that was generated
composer card delete -c admin@smarthomes
composer card import -f ./admin@smarthomes.card

# List out the network apps for this card
composer network list  -c admin@smarthomes

# ping
composer network ping -c admin@smarthomes

# list
composer network list -c admin@smarthomes

# Launch REST Server
composer-rest-server -c admin@smarthomes -n never
