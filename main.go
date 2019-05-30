package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

func main() {
	err := shim.Start(new(AssetChaincode))
	if err != nil {
		logger.Criticalf("[chaincode.main] Error starting chaincode: %s", err.Error())
	}
}