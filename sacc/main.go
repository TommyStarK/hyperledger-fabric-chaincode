package main

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	if err := shim.Start(new(SimpleAssetChaincode)); err != nil {
		fmt.Fprintf(os.Stderr, "[chaincode.main] Error starting chaincode: %s", err.Error())
	}
}
