/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

func configureStub() (*MockContext, *MockStub) {
	var nilBytes []byte

	testEeeeeeeeeegolang := new(Eeeeeeeeeegolang)
	testEeeeeeeeeegolang.Value = "set value"
	eeeeeeeeeegolangBytes, _ := json.Marshal(testEeeeeeeeeegolang)

	ms := new(MockStub)
	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
	ms.On("GetState", "missingkey").Return(nilBytes, nil)
	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
	ms.On("GetState", "eeeeeeeeeegolangkey").Return(eeeeeeeeeegolangBytes, nil)
	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

	mc := new(MockContext)
	mc.On("GetStub").Return(ms)

	return mc, ms
}

func TestEeeeeeeeeegolangExists(t *testing.T) {
	var exists bool
	var err error

	ctx, _ := configureStub()
	c := new(EeeeeeeeeegolangContract)

	exists, err = c.EeeeeeeeeegolangExists(ctx, "statebad")
	assert.EqualError(t, err, getStateError)
	assert.False(t, exists, "should return false on error")

	exists, err = c.EeeeeeeeeegolangExists(ctx, "missingkey")
	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
	assert.False(t, exists, "should return false when no value for key in world state")

	exists, err = c.EeeeeeeeeegolangExists(ctx, "existingkey")
	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
	assert.True(t, exists, "should return true when value for key in world state")
}

func TestCreateEeeeeeeeeegolang(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(EeeeeeeeeegolangContract)

	err = c.CreateEeeeeeeeeegolang(ctx, "statebad", "some value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.CreateEeeeeeeeeegolang(ctx, "existingkey", "some value")
	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

	err = c.CreateEeeeeeeeeegolang(ctx, "missingkey", "some value")
	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
}

func TestReadEeeeeeeeeegolang(t *testing.T) {
	var eeeeeeeeeegolang *Eeeeeeeeeegolang
	var err error

	ctx, _ := configureStub()
	c := new(EeeeeeeeeegolangContract)

	eeeeeeeeeegolang, err = c.ReadEeeeeeeeeegolang(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
	assert.Nil(t, eeeeeeeeeegolang, "should not return Eeeeeeeeeegolang when exists errors when reading")

	eeeeeeeeeegolang, err = c.ReadEeeeeeeeeegolang(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
	assert.Nil(t, eeeeeeeeeegolang, "should not return Eeeeeeeeeegolang when key does not exist in world state when reading")

	eeeeeeeeeegolang, err = c.ReadEeeeeeeeeegolang(ctx, "existingkey")
	assert.EqualError(t, err, "Could not unmarshal world state data to type Eeeeeeeeeegolang", "should error when data in key is not Eeeeeeeeeegolang")
	assert.Nil(t, eeeeeeeeeegolang, "should not return Eeeeeeeeeegolang when data in key is not of type Eeeeeeeeeegolang")

	eeeeeeeeeegolang, err = c.ReadEeeeeeeeeegolang(ctx, "eeeeeeeeeegolangkey")
	expectedEeeeeeeeeegolang := new(Eeeeeeeeeegolang)
	expectedEeeeeeeeeegolang.Value = "set value"
	assert.Nil(t, err, "should not return error when Eeeeeeeeeegolang exists in world state when reading")
	assert.Equal(t, expectedEeeeeeeeeegolang, eeeeeeeeeegolang, "should return deserialized Eeeeeeeeeegolang from world state")
}

func TestUpdateEeeeeeeeeegolang(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(EeeeeeeeeegolangContract)

	err = c.UpdateEeeeeeeeeegolang(ctx, "statebad", "new value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.UpdateEeeeeeeeeegolang(ctx, "missingkey", "new value")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.UpdateEeeeeeeeeegolang(ctx, "eeeeeeeeeegolangkey", "new value")
	expectedEeeeeeeeeegolang := new(Eeeeeeeeeegolang)
	expectedEeeeeeeeeegolang.Value = "new value"
	expectedEeeeeeeeeegolangBytes, _ := json.Marshal(expectedEeeeeeeeeegolang)
	assert.Nil(t, err, "should not return error when Eeeeeeeeeegolang exists in world state when updating")
	stub.AssertCalled(t, "PutState", "eeeeeeeeeegolangkey", expectedEeeeeeeeeegolangBytes)
}

func TestDeleteEeeeeeeeeegolang(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(EeeeeeeeeegolangContract)

	err = c.DeleteEeeeeeeeeegolang(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.DeleteEeeeeeeeeegolang(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

	err = c.DeleteEeeeeeeeeegolang(ctx, "eeeeeeeeeegolangkey")
	assert.Nil(t, err, "should not return error when Eeeeeeeeeegolang exists in world state when deleting")
	stub.AssertCalled(t, "DelState", "eeeeeeeeeegolangkey")
}
