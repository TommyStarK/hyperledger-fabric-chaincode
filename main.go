package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

func main() {
	logger.SetLevel(shim.LogDebug)
	err := shim.Start(new(SimpleAssetChaincode))
	if err != nil {
		logger.Criticalf("[chaincode.main] Error starting chaincode: %s", err.Error())
	}
}
