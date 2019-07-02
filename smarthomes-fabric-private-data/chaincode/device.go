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

type DeviceReading struct{
	objectType		string `json:"docType"`
	ID            	string `json:"id"`
	DeviceType 		string `json:"deviceType"`
	Data 			string `json:"data"`
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




func (t *SimpleChaincode) sendDeviceReadingPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0			1		 2
	//DeviceID	   Data		type
	if len(args[0]) <= 0 {
		return shim.Error("First argument must be a non-empty string (DeviceID)")
	}
	if len(args[1]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Temperature)")
	}

	if len(args[2]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Temperature)")
	}

	id := args[0]
	data := args[1]
	deviceType := args[2]

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

// ==================================================
// delete - remove a Device key/value pair from state
// ==================================================
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var deviceJSON DeviceReading
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
	indexName := "id~Device"
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

	fmt.Printf("- start getHistoryForDevice: %s\n", deviceID)

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
