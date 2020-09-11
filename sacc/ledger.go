package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func query(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to retrieve value of the specified key '%s': %s", args[0], err.Error())
	}

	// asset not found
	if value == nil || len(value) == 0 {
		return "", nil
	}

	return string(value), nil
}

func queryPrivateData(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetPrivateData("dummy", args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to retrieve value of the specified key '%s': %s", args[0], err.Error())
	}

	// not found
	if value == nil || len(value) == 0 {
		return "", nil
	}

	return string(value), nil
}

func store(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	var asset SimpleAsset
	if err := json.Unmarshal([]byte(args[1]), &asset); err != nil {
		return "", fmt.Errorf("Unmarshal failed: %s", err.Error())
	}

	asset.TxID = stub.GetTxID()

	b, err := json.Marshal(&asset)
	if err != nil {
		return "", fmt.Errorf("Marshal failed: %s", err.Error())
	}

	if err := stub.PutState(args[0], b); err != nil {
		return "", fmt.Errorf("Failed to put key '%s' with value (%s) into the transaction's writeset ", args[0], args[1])
	}

	return "", nil
}

func storePrivateData(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	tmap, err := stub.GetTransient()
	if err != nil {
		return "", err
	}

	for key, value := range tmap {
		if err := stub.PutPrivateData("dummy", key, value); err != nil {
			return "", fmt.Errorf("Failed to put key '%s' with value (%s) into the transaction's private writeset ", key, value)
		}
	}

	return "", nil
}

func setEvent(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting an event and its associated message")
	}

	if err := stub.SetEvent(args[0], []byte(args[1])); err != nil {
		return "", fmt.Errorf("Failed to set event: %s", args[0])
	}

	return "", nil
}
