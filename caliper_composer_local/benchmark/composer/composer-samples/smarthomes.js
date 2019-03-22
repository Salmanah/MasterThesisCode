/*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*  Perishable Goods Network
*  Performs a contractual payout for a shipment through a Transaction.
*  - Example test round
*      {
*        "label" : "perishable-network",
*        "txNumber" : [10],
*        "trim" : 0,
*        "rateControl" : [{"type": "fixed-rate", "opts": {"tps" : 10}}],
*        "arguments": {"testAssets": 10},
*        "callback" : "benchmark/composer/composer-samples/perishable-network.js"
*      }
*  - Init:
*    - Test specified number of (Importer/Grower) Participants created
*    - Test specified number of (Shipment) Assets created, belonging to a Grower Participants
*  - Run:
*    - Transactions run to perform payout upon shipment receipt
*
*/

'use strict';

const removeExisting = require('../composer-test-utils').clearAll;
const logger = require('../../../src/comm/util').getLogger('smarthomes.js');
const os = require('os');
const uuid = os.hostname() + process.pid; // UUID for client within test

module.exports.info  = 'Smarthomes Network Performance Test';

let bc;
let busNetConnection;
let testAssetNum;
let factory;
let assetId;
let testTransaction; 
const namespace = 'org.blockchain.model';

module.exports.init = async function(blockchain, context, args) {
    // Create Participants and Assets to use in main test
    bc = blockchain;
    busNetConnection = context;
    testAssetNum = args.testAssets;
    testTransaction = args.transaction; 
    assetId = 0;
     

    factory = busNetConnection.getBusinessNetwork().getFactory();

    let homeRegistry = await busNetConnection.getParticipantRegistry(namespace + '.Home');
    let ownerRegistry = await busNetConnection.getParticipantRegistry(namespace + '.Owner');
    let deviceRegistry = await busNetConnection.getParticipantRegistry(namespace + '.Device');


    let homes = new Array();
    let owners = new Array();
    let devices = new Array();
    let populated;

    switch(testTransaction){
    case 'DeviceReading':
        //Received a device reading, require Homes, owners and devices. 
            // Test specified number of Owners
        owners.push(factory.newResource(namespace, 'Owner', 'OWNER_' + uuid + '_0'));
        console.log("--------- CREATED NEW OWNER---------");
        console.log("OWNER_ID,",owners);
        // Test specified number of Homes
        for(let i = 0; i<testAssetNum; i++){
            let home = factory.newResource(namespace, 'Home', 'HOME_' + uuid + '_'+i);
            let homeAddress = factory.newConcept(namespace, 'Address');
            homeAddress.street = 'Problemveien '+i;
            homeAddress.city ='Oslo';
            home.address = homeAddress;
            console.log("--------- CREATED NEW HOME---------");
            console.log("HOME_ID,",home.homeId);
            home.homeOwner = factory.newRelationship(namespace,'Owner','OWNER_' +uuid + '_0'); 
            homes.push(home);
        }
        //Test specified number of Devices
        for(let i = 0; i <testAssetNum; i++){
            let device = factory.newResource(namespace,'Device','DEVICE_'+uuid + '_' +i); 
            device.deviceLocation = 'Kitchen'; 
            device.minimumTemperature = 25.0; 
            device.maximumTemperature = 30.0
            device.nickname = uuid;
            console.log("--------- CREATED NEW DEVICE---------");
            console.log("DEVICE_ID,",device.deviceId);
            device.home = factory.newRelationship(namespace, 'Home', 'HOME_' + uuid + '_0');
            devices.push(device);  
        }
        populated = await ownerRegistry.exists(owners[0].getIdentifier());
        break;
    case 'addingHome':
        //Transaction to add a new Home
        
        let home = factory.newResource(namespace,'Home', 'HOME_'+ uuid + '_0');
        home.homeOwner = factory.newRelationship(namespace,'Owner', factory.newResource(namespace,'Owner','OWNER_'+uuid + '_'+testAssetNum)); 
        let homeAddress = factory.newConcept(namespace, 'Address');
        homeAddress.street = 'Problemveien 221';
        homeAddress.city ='Oslo';
        home.address = homeAddress;
        homes.push(home); 
        populated = await homeRegistry.exists(homes[0].getIdentifier());
        break;
    default:
        throw new Error("No test transaction specified in module.init"); 
    }

    try {
        // Conditionally add/update registries
        
        if (!populated) {
            logger.debug('Adding test assets ...');
            logger.debug(homes);
            logger.debug("-----------------------");
            logger.debug(owners);
            logger.debug("-----------------------");
            logger.debug(devices);
            logger.debug("-----------------------");

            await homeRegistry.addAll(homes);
            await ownerRegistry.addAll(owners);
            await deviceRegistry.addAll(devices);
            logger.debug('Asset addition complete ...');
        } else {
            logger.debug('Updating test assets ...');
            logger.debug(homes);
            logger.debug("-----------------------");
            logger.debug(owners);
            logger.debug("-----------------------");
            logger.debug(devices);
            logger.debug("-----------------------");
            await removeExisting(homeRegistry, 'HOME_' + uuid);
            await removeExisting(ownerRegistry, 'OWNER_' + uuid);
            await removeExisting(deviceRegistry, 'DEVICE_' + uuid);
            await homeRegistry.addAll(homes);
            await ownerRegistry.addAll(owners);
            await deviceRegistry.addAll(devices);
            logger.debug('Asset update complete ...');
        }
    } catch (error) {
        logger.error('error in test init(): ', error);
        return Promise.reject(error);
    }
};

module.exports.run = function() {
    let transaction; 
    switch(testTransaction){
    case 'DeviceReading': {
        transaction = factory.newTransaction(namespace, 'DeviceReading');
        transaction.device = factory.newRelationship(namespace, 'Device', 'DEVICE_' + uuid + '_' +  assetId++);
        transaction.temperature = 40.0; 
        let now = new Date(); 
        now.setDate(now.getDate());
        transaction.readingDate = now; 
        transaction.readingTime = "12:00";
        break;
    }
    case 'addingHome': {
        transaction = factory.newTransaction(namespace, 'AddingHome'); 
        transaction.home = factory.newRelationship(namespace, 'Home', 'HOME_'+uuid+'_');
        break; 
    }
    default: {
        throw new Error("No valid test transaction specified in module.run");
    }
    }
    
    logger.debug("SENDING TRANSACTION DeviceReading");
    logger.debug(transaction); 
    logger.debug("---------------------------------------");
    return bc.bcObj.submitTransaction(busNetConnection, transaction);
};

module.exports.end = function() {
    return Promise.resolve(true);
};