package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func TestQueryWithMissingArgs(t *testing.T) {
	result := mstub.MockInvoke("test-query-with-missing-args", [][]byte{[]byte("query")})
	if result.Status != shim.ERROR {
		t.Log("chaincode invoke should have failed because 'query' invoke is expecting one arg: key and we provided none")
		t.Fail()
	}
}

func TestStoreWithMissingArgs(t *testing.T) {
	result := mstub.MockInvoke("test-store-with-missing-args", [][]byte{[]byte("store"), []byte("testmissingarg")})
	if result.Status != shim.ERROR {
		t.Log("chaincode invoke should have failed because 'store' invoke is expecting two args: key/value and we only provided one")
		t.Fail()
	}
}

func TestStoreThenQuery(t *testing.T) {
	resultStore := mstub.MockInvoke("test-store", [][]byte{[]byte("store"), []byte("test1"), []byte(`{"content": "this is a content test"}`)})
	if resultStore.Status != shim.OK {
		t.Log("chaincode invoke 'store' should have succeed")
		t.Fail()
	}

	resultQuery := mstub.MockInvoke("test-query", [][]byte{[]byte("query"), []byte("test1")})
	if resultQuery.Status != shim.OK {
		t.Log("chaincode invoke 'get' should have succeed")
		t.Fail()
	}

	var asset = &SimpleAsset{}
	if err := json.Unmarshal(resultQuery.Payload, asset); err != nil {
		t.Fatal(err)
	}

	if asset.Content != "this is a content test" {
		t.Log(`asset content should be a string with value "this is a content test"`)
		t.Fail()
	}

	if asset.TxID != "test-store" {
		t.Log(`asset transaction ID should be a string with value "store-asset1"`)
		t.Fail()
	}
}
