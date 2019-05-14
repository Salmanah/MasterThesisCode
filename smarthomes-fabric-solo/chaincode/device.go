/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
//Struct for a Home
type Home struct {
	ObjectType    string `json:"Home"` //docType is used to distinguish the various types of objects in state database
	ID            string `json:"id"`      //the fieldtags are needed to keep case from bouncing around
	Owner         Owner `json:"owner"`
	Address		  Address `json:"address"`
	Devices		  []Device	`json:"devices"`	
}

type Owner struct{
	ObjectType    string `json:"Owner"` //docType is used to distinguish the various types of objects in state database
	ID            string `json:"id"`
	Firstname	  string `json:"firstname"`
	Lastname	  string `json:"lastname"`	 
}

//Struct for a device
type Device struct{
	ObjectType    string `json:"Device"` //docType is used to distinguish the various types of objects in state database
	ID            string `json:"id"`      //the fieldtags are needed to keep case from bouncing around
}

type DeviceReading struct{
	ID            	string `json:"id"`
	Temperature 	string `json:"deviceReading"`
	MinimumTemp	  	string `json:"minimumTemp"`
	MaximumTemp	  	string `json:"maximumTemp"`
}

//Struct for an address
type Address struct{
	Street		string `json:"street"`
	City 		string `json:"city"` 
}


// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initDevicePrivate" { //create a new Device
		return t.initDevicePrivate(stub, args)
	}else if function == "delete" { //delete a marble
		return t.delete(stub, args)
	}else if function == "initDevice" { //delete a marble
		return t.initDevice(stub, args)
	}else if function == "readDevice" { //read a Device
		return t.readDevice(stub, args)
	}else if function == "initOwnerPrivate" { //read a Device
		return t.initOwnerPrivate(stub, args)
	}else if function == "initHomePrivate" { //read a Device
		return t.initHomePrivate(stub, args)
	}else if function == "readDevicePrivate" { //read a Device
		return t.readDevicePrivate(stub, args)
	}else if function == "sendDeviceReadingPrivate" {
		return t.sendDeviceReadingPrivate(stub, args)	
	}else if function == "sendDeviceReading" {
		return t.sendDeviceReading(stub, args)	
	}else if function == "initEnvironmentPrivate" {
		return t.initEnvironmentPrivate(stub)	
	}else if function == "getHistoryForDevice" { //get history of values for a Device
		return t.getHistoryForDevice(stub, args)
	}else if function == "readTemperature" { //get history of values for a Device
		return t.readTemperature(stub, args)
	}else if function == "readTemperaturePrivate" { //get history of values for a Device
		return t.readTemperature(stub, args)
	}else if function == "initEnvironment" { //get history of values for a Device
		return t.initEnvironment(stub)
	}else if function == "readHomeDevice" { //get history of values for a Device
		return t.readHomeDevice(stub,args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}


// ============================================================
// initMarble - create a new Device, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0             1  	      
	// "DeviceID",  "home.id" 
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init device")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	id := args[0]
	homeID := args[1]
	
	// ==== Check if device already exists ====
	deviceAsBytes, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Failed to get device: " + err.Error())
	} else if deviceAsBytes != nil {
		fmt.Println("This device already exists: " + id)
		return shim.Error("This device already exists: " + id)
	}

	//Check if home exists
	homeAsBytes, err := stub.GetState(homeID)
	if err != nil{
		return shim.Error("Failed to get home error - "+homeID)
	}

	if homeAsBytes == nil{
		fmt.Println("Home does not exist - " + homeID)
		fmt.Println(homeAsBytes)
		return shim.Error("Home does not exist -" + homeID)
	}
	
	res:= &Home{}
	json.Unmarshal(homeAsBytes, &res)

	// ==== Create device object ====
	objectType := "Device"
	device := Device{
		ObjectType: objectType,
		ID: id,
	}

	//Add device to home

	res.Devices = append(res.Devices, device)
	homeJSONBytes, err := json.Marshal(res)

	if err != nil{
		return shim.Error("Marshaling home to add device failed in initDevice")
	}

	err = stub.PutState(res.ID,homeJSONBytes)

	if err != nil{
		return shim.Error("Failed to add device to home in initDevice")
	}

	// ===== Marshal device object ======
	deviceJSONasBytes, err := json.Marshal(device)

	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save Device to state ===
	err = stub.PutState(id, deviceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}


	// ==== Device saved. Return success ====
	fmt.Println("- end init device")
	return shim.Success(deviceJSONasBytes)
}



// ============================================================
// initOwner - create a new owner
// ============================================================
func (t *SimpleChaincode) initOwnerPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// 0		1			2
	// id	firstname	lastname 
	if len(args) != 3{
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	fmt.Println("- start init owner")

	//Input sanitization
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3nd argument must be a non-empty string")
	}

	id := args[0]
	firstname := args[1]
	lastname := args[2]

	ownerAsBytes, err := stub.GetPrivateData("collectionSmarthomes",id)
	if err != nil{
		return shim.Error("Failed to get owner - "+id)
	}

	if ownerAsBytes != nil {
		return shim.Error("Owner does exist!" )
	}

	newOwner := Owner{
		ObjectType: "Owner",
		ID : id,
		Firstname: firstname,
		Lastname: lastname,
	}

	ownerJSONasBytes, err1 := json.Marshal(newOwner)

	err1 = stub.PutPrivateData("collectionSmarthomes",id, ownerJSONasBytes)
	if err1 != nil{
		return shim.Error(err1.Error())
	}

	fmt.Println("- end init Owner")
	return shim.Success(nil)
}

// ============================================================
// initHome - create a new Home, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initHomePrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 2{
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	fmt.Println("- start init home")
	if len(args[0]) <= 0{
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0{
		return shim.Error("2nd argument must be a non-empty string")
	}

	id:= args[0]
	ownerID := args[1]
	devices := []Device{}

	//Check if owner exists
	ownerAsBytes, err := stub.GetPrivateData("collectionSmarthomes",ownerID)
	if err != nil {
		return shim.Error("Failed to get owner - "+ ownerID)
	}

	if ownerAsBytes != nil{
		return shim.Error("Owner does exist")
	}

	newOwner := Owner{
		ObjectType :"Owner",
		ID: "OWNER_4",
		Firstname: "New guy",
		Lastname: "Derek",
	}

	//Check if home exists
	homeAsBytes, err3 := stub.GetPrivateData("collectionSmarthomes",id)
	if err3 != nil {
		return shim.Error("Failed to get home")
	}

	if homeAsBytes != nil{
		return shim.Error("Home does exist")
	}

	newHome := &Home{
		ObjectType: "Home",
		ID: id,
		Owner: newOwner,
		Address: Address{
			Street: "Problemveien 35b",
			City: "Oslo",
		},
		Devices: devices,
	}

	homeJSONasBytes, err4 := json.Marshal(newHome)

	err4 = stub.PutPrivateData("collectionSmarthomes",id,homeJSONasBytes)
	if err4 != nil{
		return shim.Error(err4.Error())
	}

	fmt.Println(". end init home")
	return shim.Success(nil)
}


// ============================================================
// initMarble - create a new Device, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initDevicePrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0              1   
	// "DeviceID",	"home.id"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init device")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	id := args[0]
	homeID := args[2]

	//Check if home exists
	homeAsBytes, err1 := stub.GetPrivateData("collectionSmarthomes",homeID)

	if err1 != nil {
		return shim.Error("Failed to get home - "+ homeID)
	}

	if homeAsBytes == nil{
		fmt.Println("Home does not exist - " + homeID)
		fmt.Println(homeAsBytes)
		return shim.Error("Home does not exist -" + homeID)
	}

	res := &Home{}
	json.Unmarshal(homeAsBytes,&res)

	// ==== Check if device already exists ====
	deviceAsBytes, err := loadDevice(stub,id)
	if err == nil {
		fmt.Println("This device already exists - "+id)
		fmt.Println(deviceAsBytes)
		return shim.Error("This device already exists" + id)
	}

	// ==== Create device object and marshal to JSON ====
	device := Device{
		ObjectType :"Device",
		ID: id,
	}
	deviceJSONasBytes, err3 := json.Marshal(device)

	// === Save Device to state ===
	err3 = stub.PutPrivateData("collectionSmarthomes",id, deviceJSONasBytes)
	if err3 != nil {
		return shim.Error(err3.Error())
	}

	
	//add device to home

	res.Devices = append(res.Devices,device)
	homeJSONasBytes, err4 := json.Marshal(res)

	err4 = stub.PutPrivateData("collectionSmarthomes",res.ID,homeJSONasBytes)

	 if err4 != nil{
		 return shim.Error(err4.Error())
	 }

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init device")
	return shim.Success(nil)
}

func (t *SimpleChaincode) sendDeviceReadingPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0			1
	//DeviceID	Temperature
	if len(args[0]) <= 0 {
		return shim.Error("First argument must be a non-empty string (DeviceID)")
	}
	if len(args[1]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Temperature)")
	}

	id := args[0]
	temperature := args[1]
	
	deviceAsBytes, err := loadDevice(stub, id)
	if err != nil{
		return shim.Error("Failed to get device:")
	}

	if deviceAsBytes == nil{
		return shim.Error("SendDeviceReadingPrivate device does not exist")
	}
	
	readingAsBytes,err := stub.GetPrivateData("CollectionSmarthomesPrivate",id+"-Reading") 
	if err != nil{
		return shim.Error("Failed to get reading:")
	}

	if readingAsBytes == nil{
		reading := DeviceReading{
			ID: id,
			Temperature:temperature,
			MaximumTemp: "100",
			MinimumTemp: "0",
		}

		readingJSONBytes, err := json.Marshal(reading)

		if err != nil{
			return shim.Error("Marshaling private readings failed - "+id+"-Reading")
		}

		err = stub.PutPrivateData("CollectionSmarthomesPrivate",id+"-Reading",readingJSONBytes)

		if err != nil{
			shim.Error(err.Error())
		}
		return shim.Success([]byte("Asset received new temperature stats"))
	}

	res:= DeviceReading{}
	json.Unmarshal(readingAsBytes,&res)

	res.ID = id+"-Reading"
	res.Temperature = temperature

	readingJSONBytes, err := json.Marshal(res)

	if err != nil{
		shim.Error("Marshaling private readings failed step 2 -"+id+"-Reading")
	}

	err = stub.PutPrivateData("CollectionSmarthomesPrivate",id+"-Reading",readingJSONBytes)

	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Asset modified, private reading added"))
}

func (t *SimpleChaincode) sendDeviceReading(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	
	// 0			1
	//DeviceID	Temperature
	if len(args[0]) <= 0 {
		return shim.Error("First argument must be a non-empty string (DeviceID)")
	}
	if len(args[1]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Temperature)")
	}
	
	id:= args[0]
	temperature := args[1]
	
	deviceAsBytes, err := stub.GetState(id)
	if err != nil{
		return shim.Error("Failed to get device:")
	}

	if deviceAsBytes == nil{
		return shim.Error("SendDeviceReading device does not exist "+id)
	}

	readingAsBytes, err := stub.GetState(id+"-Reading")
	if err != nil{
		return shim.Error("Failed to get deviceReading - "+id)
	}

	if readingAsBytes == nil{
		reading := DeviceReading{
			ID: id+"-Reading",
			Temperature: temperature,
		}

		readingJSONBytes, err := json.Marshal(reading)
		if err != nil{
			return shim.Error("Marshaling readings failed - "+id)
		}
		err = stub.PutState(id+"-Reading",readingJSONBytes)

		if err != nil{
			shim.Error("Failed to add readings to the blockchain - "+id+"-Reading")
		}

	}

	res := DeviceReading{}
	json.Unmarshal(readingAsBytes, &res)

	res.ID = id+"-Reading"
	res.Temperature = temperature

	readingJSONBytes, err := json.Marshal(res)

	if err != nil{
		shim.Error("Marshaling readings failed step 2 -"+id+"-Reading")
	}

	err = stub.PutState(res.ID, readingJSONBytes)
	if err != nil{
		return shim.Error("Failed to save device")
	}
	
	return shim.Success([]byte("Asset modified, new temperature"))
}

// ===============================================
// initEnvironmentPrivate - creates 2 Homes, 2 devices and 2 Owners with private data
// ===============================================

func (t *SimpleChaincode) initEnvironmentPrivate(stub shim.ChaincodeStubInterface) pb.Response {
	var err error
	fmt.Println("starting initEnvironment")

	owner1 := &Owner{
		ObjectType:"home_owner",
		ID: "OWNER_001",
		Firstname: "Tony",
		Lastname: "Stark",	
	}

	owner2 := &Owner{
		ObjectType:"home_owner",
		ID: "OWNER_002",
		Firstname: "Bruce",
		Lastname: "Wayne",	
	}

	address1 := &Address{
		Street: "Problemveien 21b",
		City: "Oslo",
	}

	address2 := &Address{
		Street: "Problemveien 21a",
		City: "Oslo",
	}

	device1 := &Device{
		ObjectType: "Device",
		ID: "DEVICE_001",
	
	}

	reading1 :=&DeviceReading{
		ID: "DEVICE_001-Reading",
		Temperature: "30",
		MinimumTemp: "15",
		MaximumTemp: "40",
	}


	device2 := &Device{
		ObjectType: "Device",
		ID: "DEVICE_002",
	}

	reading2 :=&DeviceReading{
		ID: "DEVICE_002-Reading",
		Temperature: "40",
		MinimumTemp: "25",
		MaximumTemp: "60",
	}

	home1 := &Home{
		ObjectType: "Home",
		ID: "HOME_001",
		Owner: Owner{
			ID: owner1.ID,
			Firstname:owner1.Firstname,
			Lastname: owner1.Lastname,
		},
		Address: Address{
			Street: address1.Street,
			City: address1.City,
		},
		Devices: []Device{
			Device{
				ID: device1.ID,
			},
		},
	}

	home2 := &Home{
		ObjectType: "Home",
		ID: "HOME_002",
		Owner: Owner{
			ID: owner2.ID,
			Firstname:owner2.Firstname,
			Lastname: owner2.Lastname,
		},
		Address: Address{
			Street: address2.Street,
			City: address2.City,
		},
		Devices: []Device{
			Device{
				ID: device2.ID,
			},
		},
	}


	ownerBytes, err :=json.Marshal(owner1)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = stub.PutPrivateData("collectionSmarthomes",owner1.ID, ownerBytes)

	if err != nil {
		fmt.Println("Could not store owner 1")
		return shim.Error(err.Error())
	}

	owner2Bytes, err2 :=json.Marshal(owner2)

	if err != nil{
		return shim.Error(err2.Error())
	}
	err2 = stub.PutPrivateData("collectionSmarthomes",owner2.ID, owner2Bytes)

	if err2 != nil {
		fmt.Println("Could not store owner 2")
		return shim.Error(err2.Error())
	}

	deviceBytes, err3 :=json.Marshal(device1)

	if err3 != nil{
		return shim.Error(err3.Error())
	}
	err3 = stub.PutPrivateData("collectionSmarthomes",device1.ID, deviceBytes)

	if err3 != nil {
		fmt.Println("Could not store device 1")
		return shim.Error(err3.Error())
	}

	readingBytes, err4 :=json.Marshal(reading1)

	if err4 != nil{
		return shim.Error(err4.Error())
	}
	err4 = stub.PutPrivateData("collectionSmarthomesPrivate",device1.ID, readingBytes)

	if err4 != nil {
		fmt.Println("Could not store reading 1")
		return shim.Error(err4.Error())
	}

	device2Bytes, err5 :=json.Marshal(device2)

	if err5 != nil{
		return shim.Error(err4.Error())
	}
	err5 = stub.PutPrivateData("collectionSmarthomes",device2.ID, device2Bytes)

	if err5 != nil {
		fmt.Println("Could not store device 2")
		return shim.Error(err5.Error())
	}

	reading2Bytes, err6 :=json.Marshal(reading2)

	if err6 != nil{
		return shim.Error(err6.Error())
	}
	err6 = stub.PutPrivateData("collectionSmarthomesPrivate",device2.ID, reading2Bytes)

	if err6 != nil {
		fmt.Println("Could not store reading 2")
		return shim.Error(err6.Error())
	}

	homeBytes, err7 :=json.Marshal(home1)

	if err7 != nil {
		return shim.Error(err7.Error())
	}
	err7 = stub.PutPrivateData("collectionSmarthomes",home1.ID, homeBytes)

	if err7 != nil {
		fmt.Println("Could not store home 1")
		return shim.Error(err7.Error())
	}

	home2Bytes, err8 :=json.Marshal(home2)

	if err8 != nil{
		return shim.Error(err8.Error())
	}
	err8 = stub.PutPrivateData("collectionSmarthomes", home2.ID, home2Bytes)

	if err8 != nil {
		fmt.Println("Could not store home 2")
		return shim.Error(err8.Error())
	}
	fmt.Printf("Environment created!")
	return shim.Success(homeBytes)
}


// ===============================================
// initEnvironment - creates 2 Homes, 2 devices and 2 Owners 
// ===============================================


func (t *SimpleChaincode) initEnvironment(stub shim.ChaincodeStubInterface) pb.Response {
	var err error
	fmt.Println("starting initEnvironment")

	owner1 := &Owner{
		ObjectType:"home_owner",
		ID: "OWNER_001",
		Firstname: "Tony",
		Lastname: "Stark",	
	}

	owner2 := &Owner{
		ObjectType:"home_owner",
		ID: "OWNER_002",
		Firstname: "Bruce",
		Lastname: "Wayne",	
	}

	address1 := &Address{
		Street: "Problemveien 21b",
		City: "Oslo",
	}

	address2 := &Address{
		Street: "Problemveien 21a",
		City: "Oslo",
	}

	device1 := &Device{
		ObjectType: "Device",
		ID: "DEVICE_001",
	
	}

	reading1 :=&DeviceReading{
		ID: "DEVICE_001-Reading",
		Temperature: "30",
		MinimumTemp: "15",
		MaximumTemp: "40",
	}


	device2 := &Device{
		ObjectType: "Device",
		ID: "DEVICE_002",
	}

	reading2 :=&DeviceReading{
		ID: "DEVICE_002-Reading",
		Temperature: "40",
		MinimumTemp: "25",
		MaximumTemp: "60",
	}

	home1 := &Home{
		ObjectType: "Home",
		ID: "HOME_001",
		Owner: Owner{
			ID: owner1.ID,
			Firstname:owner1.Firstname,
			Lastname: owner1.Lastname,
		},
		Address: Address{
			Street: address1.Street,
			City: address1.City,
		},
		Devices: []Device{
			Device{
				ID: device1.ID,
			},
		},
	}

	home2 := &Home{
		ObjectType: "Home",
		ID: "HOME_002",
		Owner: Owner{
			ID: owner2.ID,
			Firstname:owner2.Firstname,
			Lastname: owner2.Lastname,
		},
		Address: Address{
			Street: address2.Street,
			City: address2.City,
		},
		Devices: []Device{
			Device{
				ID: device2.ID,
			},
		},
	}


	ownerBytes, err :=json.Marshal(owner1)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(owner1.ID, ownerBytes)

	if err != nil {
		fmt.Println("Could not store owner 1")
		return shim.Error(err.Error())
	}

	owner2Bytes, err2 :=json.Marshal(owner2)

	if err != nil{
		return shim.Error(err2.Error())
	}
	err2 = stub.PutState(owner2.ID, owner2Bytes)

	if err2 != nil {
		fmt.Println("Could not store owner 2")
		return shim.Error(err2.Error())
	}

	deviceBytes, err3 :=json.Marshal(device1)

	if err3 != nil{
		return shim.Error(err3.Error())
	}
	err3 = stub.PutState(device1.ID, deviceBytes)

	if err3 != nil {
		fmt.Println("Could not store device 1")
		return shim.Error(err3.Error())
	}

	readingBytes, err4 :=json.Marshal(reading1)

	if err4 != nil{
		return shim.Error(err4.Error())
	}
	err4 = stub.PutState(reading1.ID, readingBytes)

	if err4 != nil {
		fmt.Println("Could not store reading 1")
		return shim.Error(err4.Error())
	}

	device2Bytes, err5 :=json.Marshal(device2)

	if err5 != nil{
		return shim.Error(err4.Error())
	}
	err5 = stub.PutState(device2.ID, device2Bytes)

	if err5 != nil {
		fmt.Println("Could not store device 2")
		return shim.Error(err5.Error())
	}

	reading2Bytes, err6 :=json.Marshal(reading2)

	if err6 != nil{
		return shim.Error(err6.Error())
	}
	err6 = stub.PutState(reading2.ID, reading2Bytes)

	if err6 != nil {
		fmt.Println("Could not store reading 2")
		return shim.Error(err6.Error())
	}

	homeBytes, err7 :=json.Marshal(home1)

	if err7 != nil {
		return shim.Error(err7.Error())
	}
	err7 = stub.PutState(home1.ID, homeBytes)

	if err7 != nil {
		fmt.Println("Could not store home 1")
		return shim.Error(err7.Error())
	}

	home2Bytes, err8 :=json.Marshal(home2)

	if err8 != nil{
		return shim.Error(err8.Error())
	}
	err8 = stub.PutState(home2.ID, home2Bytes)

	if err8 != nil {
		fmt.Println("Could not store home 2")
		return shim.Error(err8.Error())
	}
	fmt.Printf("Environment created!")
	return shim.Success(homeBytes)
}


// ===============================================
// readDevice - read a device from chaincode state
// ===============================================
func (t *SimpleChaincode) readDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]
	valAsbytes, err := stub.GetState(id) //get the Device from chaincode state
	
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Device does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// readDevice - read a device from chaincode state
// ===============================================
func (t *SimpleChaincode) readDevicePrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]
	valAsbytes, err := stub.GetPrivateData("collectionSmarthomesPrivate",id) //get the Device from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Device does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// readTemperature - read a device temperature stats
// ===============================================
func (t *SimpleChaincode) readTemperature(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	// 0
	// readingID

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]+"-Reading"
	valAsbytes, err := stub.GetState(id) //get the Temperature from chaincode state
	
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Device does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// readTemperaturePrivate - read a device temperature from chaincode state
// ===============================================
func (t *SimpleChaincode) readTemperaturePrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]+"-Reading"
	valAsbytes, err := stub.GetPrivateData("collectionSmarthomesPrivate",id) //get the Device from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Device does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}


// ===============================================
// readHomeDevice - read all devices of a home
// ===============================================
func (t *SimpleChaincode) readHomeDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the Home to query")
	}

	id = args[0]
	valAsbytes, err := stub.GetState(id) //get the Home from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Home does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	res := &Home{}
	json.Unmarshal(valAsbytes,&res)

	fmt.Println("%#v",res.Devices)

	return shim.Success(valAsbytes)
}





// ==================================================
// delete - remove a Device key/value pair from state
// ==================================================
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var deviceJSON Device
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	deviceID := args[0]

	// to maintain the color~name index, we need to read the Device first and get its color
	valAsbytes, err := stub.GetState(deviceID) //get the device from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + deviceID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + deviceID + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &deviceJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + deviceID + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(deviceID) //remove the Device from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "id~owner"
	colorNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{deviceJSON.ID})
	if err != nil {
		return shim.Error(err.Error())
	}

	//  Delete index entry to state.
	err = stub.DelState(colorNameIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}



func (t *SimpleChaincode) getHistoryForDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	deviceID := args[0]

	fmt.Printf("- start getHistoryForMarble: %s\n", deviceID)

	resultsIterator, err := stub.GetHistoryForKey(deviceID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the Device
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON device)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForDevice returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}



func loadDevice(stub shim.ChaincodeStubInterface, id string) (*Device, error){
	deviceID := id
	deviceBytes, err := stub.GetPrivateData("collectionSmarthomes",deviceID)
	if err != nil{
		return nil, err
	}
	res :=Device{}
	err = json.Unmarshal(deviceBytes, &res)
	if err != nil{
		return nil, err
	}
	return &res, nil
}

func loadReading(stub shim.ChaincodeStubInterface, id string) (*DeviceReading, error){
	readingID := id
	readingBytes, err := stub.GetPrivateData("collectionSmarthomesPrivate",readingID)
	if err != nil{
		return nil, err
	}
	res :=DeviceReading{}
	err = json.Unmarshal(readingBytes, &res)
	if err != nil{
		return nil, err
	}
	return &res, nil
}

func saveReading(stub shim.ChaincodeStubInterface, reading *DeviceReading) error{
	readingBytes, err := json.Marshal(reading)
	if err != nil {
		return err
	}
	id := reading.ID
	return stub.PutPrivateData("collectionSmarthomesPrivate",id, readingBytes)
}

func saveDevice(stub shim.ChaincodeStubInterface, device *Device) error {
	deviceBytes, err := json.Marshal(device)
	if err != nil {
		return err
	}
	id := device.ID
	return stub.PutPrivateData("collectionSmarthomesPrivate",id, deviceBytes)
}