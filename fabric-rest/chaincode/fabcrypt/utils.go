package fabcrypt

import(
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	"github.com/pkg/errors"
)

//Function that retrieves the value associated to the key,
//Decrypts it with the supplied entity and returns the result of the decryption

func getStateAndDecrypt(stub shim.ChaincodeStubInterface, ent entities.Encrypter, key string)([]byte, error){
	//Retrieve the ciphertext from the ledger
	ciphertext, err := stub.GetState(key)

	if err != nil{
		return nil, err
	}

	if len(ciphertext) == 0{
		return nil, errors.New("No ciphertext to decrypt")
	}

	return ent.Decrypt(ciphertext)

}

// Function that encrypts the supplied value using the supplied entity
//and puts it to the ledger associated to the supplied KVS key

func encryptAndPutState(stub shim.ChaincodeStubInterface, ent entities.Encrypter, key string, value []byte) error {
	//Use the entity to encrypt the value

	ciphertext, err := ent.Encrypt(value)

	if err != nil{
		return err
	}

	return stub.PutState(key,ciphertext)
}