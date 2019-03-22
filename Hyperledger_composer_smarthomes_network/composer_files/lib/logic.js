
'use strict';
 

/***
 * A temperature measurement has been sent from a device
 * @param {org.blockchain.model.DeviceReading} readings - the deviceReading transaction
 * @transaction
 */

 async function readingReceived(readings){
    const device = readings.device;
    const readingTime = readings.readingTime;
    const factory = getFactory(); 
     
    console.log("Temperature received at: "+readingTime);
    console.log("From Device: "+device.deviceId); 

    if(device.deviceReadings){
        device.deviceReadings.push(readings); 
    }else{
        device.deviceReadings = [readings]; 
    }

    if(readings.temperature < device.minimumTemperature || readings.temperature > device.maximumTemperature){
      var temperatureEvent = factory.newEvent("org.blockchain.model","TemperatureThresholdEvent");
      temperatureEvent.device = device; 
      temperatureEvent.message = "Temperature has reached threshold! "+device.deviceId+" has a temperature of "+readings.temperature;
      temperatureEvent.readingTime = reading.readingTime; 
      emit(temperatureEvent); 
    }
    //add the temperature reading to the device
    const deviceRegistry = await getParticipantRegistry('org.blockchain.model.Device'); 
    await deviceRegistry.update(device);
 }


 /**
  * A transaction to add home
  *@param {org.blockchain.model.AddingHome} addingHome - the home that has been added 
  *@transaction  
  */

 async function addingHome(addingHome){
     console.log("Adding home", addingHome.homeId); 
   
     const factory = getFactory(); 
     const namespace = "org.blockchain.model"; 
     
     const newHome = factory.newResource(namespace, 'Home', addingHome.homeId); 
     newHome.homeOwner = addingHome.homeOwner; 
     newHome.address = addingHome.address; 

     const registry = await getParticipantRegistry(namespace,'.Home'); 
     await registry.add(newHome); 
     const homeAddedEvent = factory.newEvent(namespace, "HomeAdded");
     homeAddedEvent.homeOwner = newHome.homeOwner; 
     homeAddedEvent.home = newHome; 
     homeAddedEvent.message = "HOME SUCCESFULLY ADDED"; 
     emit(homeAddedEvent); 
    
 }



/**
 * Initialize some test assets and participants useful for running a demo.
 * @param {org.blockchain.model.SetupDemo} setupDemo - the SetupDemo transaction
 * @transaction
 */
async function instantiateModelForTesting(setupDemo) {  // eslint-disable-line no-unused-vars

    const factory = getFactory();
    const NS = 'org.blockchain.model';
    
    // create the first home
    const firsthome = factory.newResource(NS, 'Home', 'Home_001');
    const firsthomeAddress = factory.newConcept(NS, 'Address');
  
    //Create the owner of the first house
    const firstowner = factory.newResource(NS,'Owner','Owner_001'); 
    firstowner.firstname = 'Salman'; 
    firstowner.lastname = 'Ahmed'; 
  
    //create the relationship between the house and the owner
    firsthome.homeOwner = factory.newRelationship(NS,"Owner",firstowner.personId);
    
    firsthomeAddress.street = 'Problemveien 21';
    firsthomeAddress.city = 'Oslo'; 
    firsthome.address = firsthomeAddress;
  
    var firstHomeEvent = factory.newEvent(NS,"HomeAdded"); 
    firstHomeEvent.homeOwner = firstowner; 
    firstHomeEvent.home = firsthome;
    firstHomeEvent.message = "Home added to network "+firsthome.personId;  
  
    
    // create the second home
    const secondhome = factory.newResource(NS, 'Home', 'Home_002');
    const secondhomeAddress = factory.newConcept(NS, 'Address');
  
    //create the owner of the second house
    const secondowner = factory.newResource(NS,'Owner','Owner_002'); 
    secondowner.firstname = 'John'; 
    secondowner.lastname = 'Wick'; 
    
    //create the relationship between the house and the owner
    secondhome.homeOwner = factory.newRelationship(NS,"Owner",secondowner.personId);
  
    secondhomeAddress.street = 'Problemveien 22';
    secondhomeAddress.city = 'Oslo'; 
    secondhome.address = secondhomeAddress;
  
    var secondHomeEvent = factory.newEvent(NS,"HomeAdded"); 
    secondHomeEvent.homeOwner = secondowner; 
    secondHomeEvent.home = secondhome;
    secondHomeEvent.message = "Home added to network "+secondhome.homeId;  
  
  
    const thirdhome = factory.newResource(NS, 'Home', 'Home_003');
    const thirdhomeAddress = factory.newConcept(NS, 'Address');
  
    //create the owner of the third house
    const thirdowner = factory.newResource(NS,'Owner','Owner_003'); 
    thirdowner.firstname = 'Frank'; 
    thirdowner.lastname = 'Sinatra'; 
  
     //create the relationship between the house and the owner
     thirdhome.homeOwner = factory.newRelationship(NS,"Owner",thirdowner.personId);
    
    thirdhomeAddress.street = 'Problemveien 23';
    thirdhomeAddress.city = 'Oslo'; 
    thirdhome.address = thirdhomeAddress;
  
    var thirdHomeEvent = factory.newEvent(NS,"HomeAdded"); 
    thirdHomeEvent.homeOwner = thirdowner; 
    thirdHomeEvent.home = thirdhome;
    thirdHomeEvent.message = "Home added to network "+thirdhome.homeId;  
  
  //Create Salmans device
    const first_device = factory.newResource(NS,'Device','DEVICE_001'); 	
    first_device.home = factory.newRelationship(NS,'Home', firsthome.homeId); 
    first_device.deviceLocation = "first_floor";  
    first_device.transmissionInterval = 4000;
    first_device.nickname = firsthome.homeOwner.firstname+"_fridge";
  
  //Create John device
    const second_device = factory.newResource(NS,'Device','DEVICE_002'); 	
    second_device.home = factory.newRelationship(NS,'Home', secondhome.homeId); 
    second_device.deviceLocation = "second_floor";  
   
    second_device.transmissionInterval = 6000; 
    second_device.nickname = secondhome.homeOwner.firstname+"_iphone";
  
  //Create Frank device
    const third_device = factory.newResource(NS,'Device','DEVICE_003'); 	
    third_device.home = factory.newRelationship(NS,'Home', thirdhome.homeId); 
    third_device.deviceLocation = "basement_floor";  
    third_device.transmissionInterval = 2000; 
    third_device.nickname = thirdhome.homeOwner.firstname+"_hub";
  
     //Create relationship between Home_001 and Device_001
  	if(typeof firsthome.devices == 'undefined'){
    	firsthome.devices = new Array(); 
      	firsthome.devices[0] = first_device; 
    }else{
    	firsthome.devices.push(first_device); 
    }

      //Create relationship between Home_002 and Device_002
  	if(typeof secondhome.devices == 'undefined'){
    	secondhome.devices = new Array(); 
      	secondhome.devices[0] = second_device; 
    }else{
    	secondhome.devices.push(second_device); 
    }

      //Create relationship between Home_003 and Device_003
  	if(typeof thirdhome.devices == 'undefined'){
    	thirdhome.devices = new Array(); 
      	thirdhome.devices[0] = third_device; 
    }else{
    	thirdhome.devices.push(third_device); 
    }
    
  
    
    // add the Homes
    const homeRegistry = await getParticipantRegistry(NS + '.Home');
    await homeRegistry.addAll([firsthome,secondhome,thirdhome]);
  
    // add the devices
    const deviceRegistry = await getParticipantRegistry(NS + '.Device');
    await deviceRegistry.addAll([first_device,second_device,third_device]);

    const ownerRegistry = await getParticipantRegistry(NS + '.Owner'); 
    await ownerRegistry.addAll([firstowner,secondowner,thirdowner]);
        
  }