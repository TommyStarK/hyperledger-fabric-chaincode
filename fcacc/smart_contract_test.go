package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var (
	testcc   *SmartContract
	teststub *shimtest.MockStub
)

func setup() {
	testcc, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic("failed to create chaincode: " + err.Error())
	}

	teststub = shimtest.NewMockStub("SmartContract", testcc)
}

func TestInitChaincode(t *testing.T) {
	result := teststub.MockInit("", nil)
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestInitLedger(t *testing.T) {
	result := teststub.MockInvoke("init-ledger", [][]byte{[]byte("Init")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryLedgerAfterInit(t *testing.T) {
	result := teststub.MockInvoke("query-ledger-after-init", [][]byte{[]byte("Query"), []byte("default-asset")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}

	var asset = &SimpleAsset{}
	if err := json.Unmarshal(result.Payload, asset); err != nil {
		t.Fatal(err)
	}

	if asset.Content != "default" {
		t.Log(`asset content should be a string with value "default"`)
		t.Fail()
	}

	if asset.TxID != "init-ledger" {
		t.Log(`asset transaction ID should be a string with value "init-ledger"`)
		t.Fail()
	}
}

func TestStoreAsset1(t *testing.T) {
	result := teststub.MockInvoke("store-asset1", [][]byte{[]byte("Store"), []byte("asset1"), []byte(`{"content": "foo"}`)})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryAsset1(t *testing.T) {
	result := teststub.MockInvoke("query-ledger-asset1", [][]byte{[]byte("Query"), []byte("asset1")})
	if result.Status != shim.OK {
		t.Logf("failed to query the ledger: %s", result.GetMessage())
		t.Fail()
	}

	var asset = &SimpleAsset{}
	if err := json.Unmarshal(result.Payload, asset); err != nil {
		t.Fatal(err)
	}

	if asset.Content != "foo" {
		t.Log(`asset content should be a string with value "foo"`)
		t.Fail()
	}

	if asset.TxID != "store-asset1" {
		t.Log(`asset transaction ID should be a string with value "store-asset1"`)
		t.Fail()
	}
}

func TestUpdateAsset1(t *testing.T) {
	result := teststub.MockInvoke("update-asset1", [][]byte{[]byte("Store"), []byte("asset1"), []byte(`{"content": "bar", "txID": "store-asset1"}`)})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryUpdatedAsset1(t *testing.T) {
	result := teststub.MockInvoke("query-updated-asset1", [][]byte{[]byte("Query"), []byte("asset1")})
	if result.Status != shim.OK {
		t.Logf("failed to query the ledger: %s", result.GetMessage())
		t.Fail()
	}

	var asset = &SimpleAsset{}
	if err := json.Unmarshal(result.Payload, asset); err != nil {
		t.Fatal(err)
	}

	if asset.Content != "bar" {
		t.Log(`asset content should be a string with value "bar"`)
		t.Fail()
	}

	if asset.TxID != "update-asset1" {
		t.Log(`asset transaction ID should be a string with value "update-asset1"`)
		t.Fail()
	}
}

func TestDeleteDefaultAsset(t *testing.T) {
	result := teststub.MockInvoke("delete-default-asset", [][]byte{[]byte("Delete"), []byte("default-asset")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryUnknownAsset(t *testing.T) {
	result := teststub.MockInvoke("query-unknown-asset", [][]byte{[]byte("Query"), []byte("unknown-asset")})
	if result.Status != shim.ERROR {
		t.Log("querying an unknown asset should have failed")
		t.Fail()
	}

	if result.Message != "failed to query the ledger for specific 'key': unknown-asset" {
		t.Log(`response message should equal "failed to query the ledger for specific 'key': unknown-asset"`)
		t.Fail()
	}

	if len(result.Payload) > 0 {
		t.Log("payload response should be empty")
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
