package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestSetWithMissingArgs(t *testing.T) {
	result := mstub.MockInvoke("testsetwithmissingarg", [][]byte{[]byte("set"), []byte("testmissingarg")})

	t.Log(result.String())
	if result.GetStatus() != shim.ERROR {
		t.Log("chaincode invoke should have failed because 'set' invoke is expecting two args: key/value and we only provided one")
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	result := mstub.MockInvoke("testset", [][]byte{[]byte("set"), []byte("test1"), []byte(`{"content": "this is a content test"}`)})

	if result.GetStatus() != shim.OK {
		t.Log("chaincode invoke should have succeed")
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	result := mstub.MockInvoke("testget", [][]byte{[]byte("get"), []byte("test1")})

	t.Log(result.String())
	if result.GetStatus() != shim.OK {
		t.Log("chaincode invoke should have succeed")
		t.Fail()
	}

	t.Log(string(result.GetPayload()))
}
