package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (sc *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	var asset = &SimpleAsset{
		Content: "default",
		TxID:    ctx.GetStub().GetTxID(),
	}

	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState("default-asset", assetAsBytes); err != nil {
		return err
	}

	return nil
}

func (sc *SmartContract) Delete(ctx contractapi.TransactionContextInterface, key string) error {
	if err := ctx.GetStub().DelState(key); err != nil {
		return err
	}

	return nil
}

func (sc *SmartContract) Query(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	bytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", err
	}

	if bytes == nil || len(bytes) == 0 {
		return "", errors.New("failed to query the ledger for specific 'key': " + key)
	}

	return string(bytes), nil
}

func (sc *SmartContract) Store(ctx contractapi.TransactionContextInterface, key, stringifiedAsset string) error {
	var asset = &SimpleAsset{}
	if err := json.Unmarshal([]byte(stringifiedAsset), asset); err != nil {
		return err
	}

	asset.TxID = ctx.GetStub().GetTxID()
	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(key, assetAsBytes); err != nil {
		return err
	}

	return nil
}
