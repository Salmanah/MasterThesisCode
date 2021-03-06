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
	DataSize		int		`json:"dataSize"`
}

type DeviceData struct{
	objectType		string `json:"docType"`
	ID            	string `json:"id"`
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
	if function == "delete" { //delete a marble
		return t.delete(stub, args)
	}else if function == "readDevice" { //read a Device
		return t.readDevice(stub, args)
	}else if function == "readDevicePrivate" { //read a Device
		return t.readDevicePrivate(stub, args)
	}else if function == "sendDeviceReadingPrivate" {
		return t.sendDeviceReadingPrivate(stub, args)	
	}else if function == "getHistoryForDevice" { //get history of values for a Device
		return t.getHistoryForDevice(stub, args)
	}/*else if function == "initDevice" { //read a Device
		return t.initDevice(stub, args)
	}*/

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) initDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0			1		
	//DeviceID	   Type		
	if len(args[0]) <= 0 {
		return shim.Error("First argument must be a non-empty string (DeviceID)")
	}
	if len(args[1]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Type)")
	}


	id := args[0]
	deviceType := args[1]
	data := ""
	dataSize := len(data)

	reading := DeviceReading{
		ID: id,
		DeviceType: deviceType,
		DataSize: dataSize,
	}

	privateReading := DeviceData{
		ID: id,
		Data:data,
	}

	readingJSONBytes, err := json.Marshal(reading)

	if err != nil{
		return shim.Error("Marshaling private readings failed - "+id)
	}

	err = stub.PutPrivateData("collectionSmarthomes",id,readingJSONBytes)

	if err != nil{
		return shim.Error(err.Error())
	}

	readingJSONPrivate, err := json.Marshal(privateReading)

	err = stub.PutPrivateData("collectionSmarthomesPrivate",id,readingJSONPrivate)

	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Device intitiated"))
}
/*
func (t *SimpleChaincode) sendDeviceReadingPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0			1		 2
	//DeviceID	   Type   Data	
	if len(args[0]) <= 0 {
		return shim.Error("First argument must be a non-empty string (DeviceID)")
	}
	if len(args[1]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Data)")
	}

	if len(args[2]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Data)")
	}


	id := args[0]
	deviceType := args[1]
	data := args[2]
	dataSize := len(data)
	fmt.Println("- Start sendingReadings -", id)

	deviceAsBytes, err := stub.GetPrivateData("collectionSmarthomes",id)
	deviceAsBytesPrivate, err1 := stub.GetPrivateData("collectionSmarthomesPrivate",id)

	if err != nil {
		return shim.Error("Failed to get device: "+err.Error())
	}else if deviceAsBytes == nil {
		return shim.Error("Device does not exist")
	}

	if err1 != nil {
		return shim.Error("Failed to get device: "+err.Error())
	}else if deviceAsBytesPrivate == nil {
		return shim.Error("Private Device does not exist")
	}

	newReading := DeviceReading{}
	newPrivateReading := DeviceData{}

	err = json.Unmarshal(deviceAsBytes, &newReading)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = json.Unmarshal(deviceAsBytesPrivate, &newPrivateReading)
	if err != nil{
		return shim.Error(err.Error())
	}

	newPrivateReading.Data = data
	newReading.DataSize = dataSize
	newReading.DeviceType = deviceType


	readingJSONBytes, err := json.Marshal(newReading)

	if err != nil{
		return shim.Error("Marshaling private readings failed - "+id)
	}

	err = stub.PutPrivateData("collectionSmarthomes",id,readingJSONBytes)

	if err != nil{
		return shim.Error(err.Error())
	}

	readingJSONPrivate, err := json.Marshal(newPrivateReading)

	err = stub.PutPrivateData("collectionSmarthomesPrivate",id,readingJSONPrivate)

	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Data successfully updated"))
}*/

func (t *SimpleChaincode) sendDeviceReadingPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	
	// 0			1		 2
	//DeviceID	   data		type
	if len(args[0]) <= 0 {
		return shim.Error("First argument must be a non-empty string (DeviceID)")
	}
	if len(args[1]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Type)")
	}

	if len(args[2]) <= 0{
		return shim.Error("Third argument must be a non-empty string (Data)")
	}
	
	id:= args[0]
	deviceType := args[1]
	data := args[2]
	dataSize:= len(deviceType)
	
	reading := DeviceReading{
		objectType : "docType",
		ID: id,
		DeviceType:deviceType,
		DataSize:dataSize,
		
	}

	privateData := DeviceData{
		ID: id,
		Data: data,
	}

	readingJSONBytes, err := json.Marshal(reading)
	if err != nil{
		return shim.Error("Marshaling readings failed - "+id)
	}

	readingJSONBytesPrivate, err := json.Marshal(privateData)
	if err != nil{
		return shim.Error("Marshaling readings failed - "+id)
	}

	err = stub.PutPrivateData("collectionSmarthomes",id,readingJSONBytes)

	if err != nil{
		shim.Error("Failed to add readings to the blockchain - "+id)
	}

	err = stub.PutPrivateData("collectionSmarthomesPrivate",id,readingJSONBytesPrivate)
	
	return shim.Success([]byte("Asset modified!"))
}

// ===============================================
// readDevice - read a device from chaincode state
// ===============================================
func (t *SimpleChaincode) readDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id string

	deviceReading := DeviceReading{}

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]
	valAsbytes, err1 := stub.GetPrivateData("collectionSmarthomes",id)  //get the Device from chaincode state
	
	if err1 != nil {
		return shim.Error(err1.Error())
	}

	err2:= json.Unmarshal(valAsbytes, &deviceReading)

	if err2 != nil{
		fmt.Println("Error unmarshalling object with id: "+id)
		return shim.Error(err2.Error())
	}
	jsonReading, err3 := json.Marshal(deviceReading)

	if err3 != nil{
		return shim.Error(err3.Error())
	}

	return shim.Success(jsonReading)
}

// ===============================================
// readDevice - read a device from chaincode state
// ===============================================
func (t *SimpleChaincode) readDevicePrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id string
	privateData := DeviceData{}

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]
	valAsbytes, err1 := stub.GetPrivateData("collectionSmarthomesPrivate",id)  //get the Device from chaincode state
	
	if err1 != nil {
		return shim.Error(err1.Error())
	}

	err2:= json.Unmarshal(valAsbytes, &privateData)

	if err2 != nil{
		fmt.Println("Error unmarshalling object with id: "+id)
		return shim.Error(err2.Error())
	}
	jsonReading, err3 := json.Marshal(privateData)

	if err3 != nil{
		return shim.Error(err3.Error())
	}

	return shim.Success(jsonReading)
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
	valAsbytes, err := stub.GetPrivateData("collectionSmarthomes",deviceID)  //get the device from chaincode state
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
