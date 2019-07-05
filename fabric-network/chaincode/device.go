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

	fc "github.com/chaincode/fabcrypt"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
       
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// DeviceReading struct 
type DeviceReading struct{
	objectType		string `json:"docType"` 
	ID            	string `json:"id"`
	Type			string `json:"Type"`
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
	} else if function == "readDevice" { //read a Device
		return t.readDevice(stub, args)
	} else if function == "sendDeviceReading" {
		return t.sendDeviceReading(stub, args)	
	 }else if function == "sendDeviceReadingEncrypted" {
		return t.sendDeviceReadingEncrypted(stub, args)	
	 }else if function == " readDeviceEncrypted" {
		return t. readDeviceEncrypted(stub, args)	
	 }

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) sendDeviceReading(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	
	// 		1		2	
	//	   Type   Data 
	
	if len(args[0]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Type)")
	}
	if len(args[1]) <= 0{
		return shim.Error("Second argument must be a non-empty string (Data)")
	}
	
	id,err := cid.New(stub)
	
	if err != nil{
		return shim.Error("Getting clientID failed", err) 
	}

	type := args[0]
	temperature := args[1]
	
	//Create new deviceReading object
	reading := DeviceReading{
		objectType := "reading"
		ID: id,
		Temperature: temperature,
	}

	readingJSONBytes, err := json.Marshal(reading)
	if err != nil{
		return shim.Error("Marshaling readings failed - "+id)
	}
	err = stub.PutState(id,readingJSONBytes)

	if err != nil{
		shim.Error("Failed to add readings to the blockchain - "+id)
	}

	stub.SetEvent(name:"DeviceReading", "payload","DeviceReading successfully sent")
	return shim.Success([]byte("DeviceReading successfully sent"))
}

// ===============================================
// Sends encrypted deviceReading
// ===============================================


func (t *SimpleChaincode) sendDeviceReadingEncrypted(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

	err := fc.Encrypter(stub, id, temperature)
	
	if err != nil{
		return shim.Error(err.Error)
	}

	return shim.Success([]byte("Asset modified, new temperature"))
}

// ===============================================
// initEnvironment - creates 2 Homes, 2 devices and 2 Owners 
// ===============================================




// ===============================================
// readDeviceEncrypted - Decrypts the data and prints it
// ===============================================
func (t *SimpleChaincode) readDeviceEncrypted(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]
	key := args[0]
	valAsbytes, err := fc.Decrypter(stub,key)
	
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
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
	deviceBytes, err := stub.GetState(deviceID)
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

func saveDevice(stub shim.ChaincodeStubInterface, device *Device) error {
	deviceBytes, err := json.Marshal(device)
	if err != nil {
		return err
	}
	id := device.ID
	return stub.PutState(id, deviceBytes)
}
