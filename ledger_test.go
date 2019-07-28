package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestHistoryWithMissingArg(t *testing.T) {
	result := mstub.MockInvoke("testhistorywithmissingarg", [][]byte{[]byte("history")})

	t.Log(result.String())
	if result.GetStatus() != shim.ERROR {
		t.Log("chaincode invoke should have failed because 'history' invoke is expecting one arg: key and we provided none")
		t.Fail()
	}
}
func TestGetWithMissingArgs(t *testing.T) {
	result := mstub.MockInvoke("testgetwithmissingarg", [][]byte{[]byte("get")})

	t.Log(result.String())
	if result.GetStatus() != shim.ERROR {
		t.Log("chaincode invoke should have failed because 'get' invoke is expecting one arg: key and we provided none")
		t.Fail()
	}
}

func TestSetWithMissingArgs(t *testing.T) {
	result := mstub.MockInvoke("testsetwithmissingarg", [][]byte{[]byte("set"), []byte("testmissingarg")})

	t.Log(result.String())
	if result.GetStatus() != shim.ERROR {
		t.Log("chaincode invoke should have failed because 'set' invoke is expecting two args: key/value and we only provided one")
		t.Fail()
	}
}

func TestSetThenGet(t *testing.T) {
	resultSet := mstub.MockInvoke("testset", [][]byte{[]byte("set"), []byte("test1"), []byte(`{"content": "this is a content test"}`)})

	t.Log(resultSet.String())
	if resultSet.GetStatus() != shim.OK {
		t.Log("chaincode invoke 'set' should have succeed")
		t.Fail()
	}

	resultGet := mstub.MockInvoke("testget", [][]byte{[]byte("get"), []byte("test1")})

	t.Log(resultGet.String())
	if resultGet.GetStatus() != shim.OK {
		t.Log("chaincode invoke 'get' should have succeed")
		t.Fail()
	}

	t.Log(string(resultGet.GetPayload()))
}

func TestHistory(t *testing.T) {
	result := mstub.MockInvoke("testhistory", [][]byte{[]byte("history"), []byte("test1")})

	t.Log(result.GetStatus())
	if result.GetStatus() != shim.OK {
		t.Log("chaincode invoke 'history' should have succeed")
		t.Fail()
	}
}
