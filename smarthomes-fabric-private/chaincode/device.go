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

/*
* NOTE: This implementation is a replica of the following:
* https://github.com/hyperledger/fabric-samples/blob/release-1.1/chaincode/Devices02/node/Devices_chaincode.js
 */

// ====CHAINCODE EXECUTION SAMPLES (CLI) ==================

// ==== Invoke Devices ====
// peer chaincode invoke -C myc1 -n Devices -c '{"Args":["initMarble","Device1","blue","35","tom"]}'
// peer chaincode invoke -C myc1 -n Devices -c '{"Args":["initMarble","Device2","red","50","tom"]}'
// peer chaincode invoke -C myc1 -n Devices -c '{"Args":["initMarble","Device3","blue","70","tom"]}'
// peer chaincode invoke -C myc1 -n Devices -c '{"Args":["transferMarble","Device2","jerry"]}'
// peer chaincode invoke -C myc1 -n Devices -c '{"Args":["transferMarblesBasedOnColor","blue","jerry"]}'
// peer chaincode invoke -C myc1 -n Devices -c '{"Args":["delete","Device1"]}'

// ==== Query Devices ====
// peer chaincode query -C myc1 -n Devices -c '{"Args":["readMarble","Device1"]}'
// peer chaincode query -C myc1 -n Devices -c '{"Args":["getMarblesByRange","Device1","Device3"]}'
// peer chaincode query -C myc1 -n Devices -c '{"Args":["getHistoryForMarble","Device1"]}'

// Rich Query (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n Devices -c '{"Args":["queryMarblesByOwner","tom"]}'
//   peer chaincode query -C myc1 -n Devices -c '{"Args":["queryMarbles","{\"selector\":{\"owner\":\"tom\"}}"]}'

// INDEXES TO SUPPORT COUCHDB RICH QUERIES
//
// Indexes in CouchDB are required in order to make JSON queries efficient and are required for
// any JSON query with a sort. As of Hyperledger Fabric 1.1, indexes may be packaged alongside
// chaincode in a META-INF/statedb/couchdb/indexes directory. Each index must be defined in its own
// text file with extension *.json with the index definition formatted in JSON following the
// CouchDB index JSON syntax as documented at:
// http://docs.couchdb.org/en/2.1.1/api/database/find.html#db-index
//
// This Devices02 example chaincode demonstrates a packaged
// index which you can find in META-INF/statedb/couchdb/indexes/indexOwner.json.
// For deployment of chaincode to production environments, it is recommended
// to define any indexes alongside chaincode so that the chaincode and supporting indexes
// are deployed automatically as a unit, once the chaincode has been installed on a peer and
// instantiated on a channel. See Hyperledger Fabric documentation for more details.
//
// If you have access to the your peer's CouchDB state database in a development environment,
// you may want to iteratively test various indexes in support of your chaincode queries.  You
// can use the CouchDB Fauxton interface or a command line curl utility to create and update
// indexes. Then once you finalize an index, include the index definition alongside your
// chaincode in the META-INF/statedb/couchdb/indexes directory, for packaging and deployment
// to managed environments.
//
// In the examples below you can find index definitions that support Devices02
// chaincode queries, along with the syntax that you can use in development environments
// to create the indexes in the CouchDB Fauxton interface or a curl command line utility.
//

//Example hostname:port configurations to access CouchDB.
//
//To access CouchDB docker container from within another docker container or from vagrant environments:
// http://couchdb:5984/
//
//Inside couchdb docker container
// http://127.0.0.1:5984/

// Index for docType, owner.
// Note that docType and owner fields must be prefixed with the "data" wrapper
//
// Index definition for use with Fauxton interface
// {"index":{"fields":["data.docType","data.owner"]},"ddoc":"indexOwnerDoc", "name":"indexOwner","type":"json"}
//
// Example curl command line to define index in the CouchDB channel_chaincode database
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[\"data.docType\",\"data.owner\"]},\"name\":\"indexOwner\",\"ddoc\":\"indexOwnerDoc\",\"type\":\"json\"}" http://hostname:port/myc1_marbles/_index
//

// Index for docType, owner, size (descending order).
// Note that docType, owner and size fields must be prefixed with the "data" wrapper
//
// Index definition for use with Fauxton interface
// {"index":{"fields":[{"data.size":"desc"},{"data.docType":"desc"},{"data.owner":"desc"}]},"ddoc":"indexSizeSortDoc", "name":"indexSizeSortDesc","type":"json"}
//
// Example curl command line to define index in the CouchDB channel_chaincode database
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[{\"data.size\":\"desc\"},{\"data.docType\":\"desc\"},{\"data.owner\":\"desc\"}]},\"ddoc\":\"indexSizeSortDoc\", \"name\":\"indexSizeSortDesc\",\"type\":\"json\"}" http://hostname:port/myc1_marbles/_index

// Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n marbles -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":\"marble\",\"owner\":\"tom\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'

// Rich Query with index design doc specified only (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n marbles -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":{\"$eq\":\"marble\"},\"owner\":{\"$eq\":\"tom\"},\"size\":{\"$gt\":0}},\"fields\":[\"docType\",\"owner\",\"size\"],\"sort\":[{\"size\":\"desc\"}],\"use_index\":\"_design/indexSizeSortDoc\"}"]}'

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
//Struct for a Home
type Home struct {
	ObjectType    string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID            string `json:"id"`      //the fieldtags are needed to keep case from bouncing around
	Owner         Owner `json:"owner"`
	Address		  Address `json:"address"`
	Devices		  []Device	`json:"devices"`	
}

type Owner struct{
	ObjectType    string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID            string `json:"id"`
	Firstname	  string `json:"firstname"`
	Lastname	  string `json:"lastname"`	 
}

//Struct for a device
type Device struct{
	ObjectType    string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID            string `json:"id"`      //the fieldtags are needed to keep case from bouncing around
	Owner		  Owner	`json:"owner"`
}

type DeviceReading struct{
	ID				string `json:"id"`
	Temperature 	int	`json:"temperature"`
	MinimumTemp	  	int `json:"minimumTemp"`
	MaximumTemp	  	int `json:"maximumTemp"`
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
	if function == "initDevice" { //create a new Device
		return t.initDevice(stub, args)
	} else if function == "delete" { //delete a marble
		return t.delete(stub, args)
	} else if function == "readDevice" { //read a Device
		return t.readDevice(stub, args)
	} else if function == "sendDeviceReading" {
		return t.sendDeviceReading(stub, args)	
	}else if function == "readTemp" {
		return t.readTemp(stub,args)	
	}else if function == "initEnvironment" {
		return t.initEnvironment(stub)	
	} else if function == "queryDeviceByOwner" { //find Devices for owner X using rich query
		return t.queryDeviceByOwner(stub, args)
	} else if function == "queryDevice" { //find Devices based on an ad hoc rich query
		return t.queryDevice(stub, args)
	} else if function == "getHistoryForDevice" { //get history of values for a Device
		return t.getHistoryForDevice(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMarble - create a new Device, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0              1 	
	// "DeviceID",  "owner.id" 
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init device")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}

	id := args[0]
	ownerID := args[4]

	var owner Owner

	//Check if owner exists
	ownerAsBytes, err := stub.GetPrivateData("collectionMedium",ownerID)
	if err != nil {
		return shim.Error("Failed to get owner - "+ ownerID)
	}

	json.Unmarshal(ownerAsBytes, &owner)

	if len(owner.Firstname) == 0 {
		return shim.Error("Owner does not exist")
	}
	
	if err != nil {
		return shim.Error("2rd argument must be a numeric string")
	}


	// ==== Check if device already exists ====
	deviceAsBytes, err := stub.GetPrivateData("collectionMedium",id)
	if err != nil {
		return shim.Error("Failed to get device: " + err.Error())
	} else if deviceAsBytes != nil {
		fmt.Println("This device already exists: " + id)
		return shim.Error("This device already exists: " + id)
	}

	// ==== Create device object and marshal to JSON ====
	objectType := "device"
	device := &Device{objectType, id, owner}
	deviceJSONasBytes, err := json.Marshal(device)
	if err != nil {
		return shim.Error(err.Error())
	}
	//Alternatively, build the Device json string manually if you don't want to use struct marshalling
	//DeviceJSONasString := `{"docType":"Marble",  "name": "` + DeviceName + `", "color": "` + color + `", "size": ` + strconv.Itoa(size) + `, "owner": "` + owner + `"}`
	//DeviceJSONasBytes := []byte(str)

	// === Save Device to state ===
	err = stub.PutPrivateData("collectionMedium",id, deviceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//  ==== Index the Device to enable color-based range queries, e.g. return all blue Devices ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~color~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexName := "id~owner"
	deviceNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{device.ID, device.Owner.ID})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the Device.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(deviceNameIndexKey, value)

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init device")
	return shim.Success(nil)
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
	
	deviceReading, err := loadDeviceReading(stub, args[0])
	if err != nil{
		return shim.Error("Failed to get device reading: " + args[0])
	}

	temperature, err := strconv.Atoi(args[1])
	
	deviceReading.Temperature = temperature;
	
	err = saveReading(stub,deviceReading)
	if err != nil{
		return shim.Error("Failed to save device")
	}
	
	return shim.Success([]byte("Asset modified, new temperature"))
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
		Owner: Owner{
			ID: owner1.ID,
			Firstname: owner1.Firstname,
			Lastname: owner1.Lastname, 
		},
	}

	deviceReading1 := &DeviceReading{
		ID: "DEVICE_001",
		Temperature: 30,
		MinimumTemp: 15,
		MaximumTemp: 40,
	}


	device2 := &Device{
		ObjectType: "Device",
		ID: "DEVICE_002",
		Owner: Owner{
			ID: owner2.ID,
			Firstname: owner2.Firstname,
			Lastname: owner2.Lastname, 
		},
	}

	deviceReading2 := &DeviceReading{
		ID: "DEVICE_002",
		Temperature: 50,
		MinimumTemp: 25,
		MaximumTemp: 60,
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
				Owner: device1.Owner,
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
				Owner: device2.Owner,
			},
		},
	}

	ownerBytes, err :=json.Marshal(owner1)
	err = stub.PutPrivateData("collectionMedium",owner1.ID, ownerBytes)

	if err != nil {
		fmt.Println("Could not store owner")
		return shim.Error(err.Error())
	}

	owner2Bytes, err :=json.Marshal(owner2)
	err = stub.PutPrivateData("collectionMedium",owner2.ID, owner2Bytes)

	if err != nil {
		fmt.Println("Could not store owner")
		return shim.Error(err.Error())
	}

	deviceBytes, err :=json.Marshal(device1)
	err = stub.PutPrivateData("collectionMedium",device1.ID, deviceBytes)

	if err != nil {
		fmt.Println("Could not store device")
		return shim.Error(err.Error())
	}

	device2Bytes, err :=json.Marshal(device2)
	err = stub.PutPrivateData("collectionMedium",device2.ID, device2Bytes)

	if err != nil {
		fmt.Println("Could not store device")
		return shim.Error(err.Error())
	}

	reading1Bytes, err :=json.Marshal(deviceReading1)
	err = stub.PutPrivateData("collectionPrivate",device1.ID,reading1Bytes)
	if err != nil {
		fmt.Println("Could not store device")
		return shim.Error(err.Error())
	}

	reading2Bytes, err :=json.Marshal(deviceReading2)
	err = stub.PutPrivateData("collectionPrivate",device2.ID,reading2Bytes)
	if err != nil {
		fmt.Println("Could not store device")
		return shim.Error(err.Error())
	}

	homeBytes, err :=json.Marshal(home1)
	err = stub.PutPrivateData("collectionMedium",home1.ID, homeBytes)

	if err != nil {
		fmt.Println("Could not store home")
		return shim.Error(err.Error())
	}

	home2Bytes, err :=json.Marshal(home2)
	err = stub.PutPrivateData("collectionMedium",home2.ID, home2Bytes)

	if err != nil {
		fmt.Println("Could not store home")
		return shim.Error(err.Error())
	}
	fmt.Printf("Environment created!")
	return shim.Success(deviceBytes)
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
	valAsbytes, err := stub.GetPrivateData("collectionPrivate",id) //get the Device from chaincode state
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
// readTemp - read a device temperature from chaincode state
// ===============================================
func (t *SimpleChaincode) readTemp(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ID of the device to query")
	}

	id = args[0]
	valAsbytes, err := stub.GetPrivateData("collectionPrivate",id) //get the Device from chaincode state
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
	colorNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{deviceJSON.ID, deviceJSON.Owner.ID})
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

// =======Rich queries =========================================================================
// Two examples of rich queries are provided below (parameterized query and ad hoc query).
// Rich queries pass a query string to the state database.
// Rich queries are only supported by state database implementations
//  that support rich query (e.g. CouchDB).
// The query string is in the syntax of the underlying state database.
// With rich queries there is no guarantee that the result set hasn't changed between
//  endorsement time and commit time, aka 'phantom reads'.
// Therefore, rich queries should not be used in update transactions, unless the
// application handles the possibility of result set changes between endorsement and commit time.
// Rich queries can be used for point-in-time queries against a peer.
// ============================================================================================

// ===== Example: Parameterized rich query =================================================
// queryMarblesByOwner queries for Devices based on a passed in owner.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *SimpleChaincode) queryDeviceByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "bob"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	owner := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"device\",\"owner\":\"%s\"}}", owner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===== Example: Ad hoc rich query ========================================================
// queryMarbles uses a query string to perform a query for Devices.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the queryMarblesForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *SimpleChaincode) queryDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
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

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func loadDevice(stub shim.ChaincodeStubInterface, id string) (*Device, error){
	deviceID := id
	deviceBytes, err := stub.GetPrivateData("collectionMedium",deviceID)
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

func loadDeviceReading(stub shim.ChaincodeStubInterface, id string) (*DeviceReading, error){
	readingID := id
	readingBytes, err := stub.GetPrivateData("collectionPrivate",readingID)

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

func saveReading(stub shim.ChaincodeStubInterface, reading *DeviceReading) error {
	readingBytes, err := json.Marshal(reading)
	if err != nil {
		return err
	}
	id := reading.ID
	return stub.PutPrivateData("collectionPrivate",id, readingBytes)
}

func saveDevice(stub shim.ChaincodeStubInterface, device *Device) error {
	deviceBytes, err := json.Marshal(device)
	if err != nil {
		return err
	}
	id := device.ID
	return stub.PutPrivateData("collectionMedium",id, deviceBytes)
}