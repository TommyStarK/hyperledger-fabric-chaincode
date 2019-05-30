package main

import (
	"fmt"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func history(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	history, err := stub.GetHistoryForKey(args[0])

	if err != nil {
		return nil, fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return nil, fmt.Errorf("Asset not found: %s", args[0])
	}

	batch := make([]string, 10)
	for value.HasNext() {
		batch = append(batch, value.Next())
	}

	return []byte(strings.Join(batch, "\n")), nil
}

func get(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return nil, fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err.Error())
	}
	if value == nil {
		return nil, fmt.Errorf("Asset not found: %s", args[0])
	}

	return value, nil
}

func set(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return nil, fmt.Errorf("Failed to set asset: %s", args[0])
	}

	return nil, nil
}
