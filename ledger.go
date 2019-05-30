package main

import (
	"fmt"
	"encoding/json"
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

	if history == nil {
		return nil, fmt.Errorf("Asset not found: %s", args[0])
	}

	batch := make([]string, 8)
	for history.HasNext() {
		modif, err := history.Next()
		if err != nil {
			continue
		}
		batch = append(batch, modif.String())
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

	var asset Asset
	if err := json.Unmarshal([]byte(args[1]), &asset); err != nil {
		return nil, fmt.Errorf("Unmarshal failed: %s", err.Error())
	}

	asset.TxID = stub.GetTxID()

	b, err := json.Marshal(&asset)
	if err != nil {
		return nil, fmt.Errorf("Marshal failed: %s", err.Error())
	}

	if err := stub.PutState(args[0], b); err != nil {
		return nil, fmt.Errorf("Failed to set asset: %s", args[0])
	}

	return nil, nil
}
