/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// EeeeeeeeeegolangContract contract for managing CRUD for Eeeeeeeeeegolang
type EeeeeeeeeegolangContract struct {
	contractapi.Contract
}

// EeeeeeeeeegolangExists returns true when asset with given ID exists in world state
func (c *EeeeeeeeeegolangContract) EeeeeeeeeegolangExists(ctx contractapi.TransactionContextInterface, eeeeeeeeeegolangID string) (bool, error) {
	data, err := ctx.GetStub().GetState(eeeeeeeeeegolangID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateEeeeeeeeeegolang creates a new instance of Eeeeeeeeeegolang
func (c *EeeeeeeeeegolangContract) CreateEeeeeeeeeegolang(ctx contractapi.TransactionContextInterface, eeeeeeeeeegolangID string, value string) error {
	exists, err := c.EeeeeeeeeegolangExists(ctx, eeeeeeeeeegolangID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", eeeeeeeeeegolangID)
	}

	eeeeeeeeeegolang := new(Eeeeeeeeeegolang)
	eeeeeeeeeegolang.Value = value

	bytes, _ := json.Marshal(eeeeeeeeeegolang)

	return ctx.GetStub().PutState(eeeeeeeeeegolangID, bytes)
}

// ReadEeeeeeeeeegolang retrieves an instance of Eeeeeeeeeegolang from the world state
func (c *EeeeeeeeeegolangContract) ReadEeeeeeeeeegolang(ctx contractapi.TransactionContextInterface, eeeeeeeeeegolangID string) (*Eeeeeeeeeegolang, error) {
	exists, err := c.EeeeeeeeeegolangExists(ctx, eeeeeeeeeegolangID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", eeeeeeeeeegolangID)
	}

	bytes, _ := ctx.GetStub().GetState(eeeeeeeeeegolangID)

	eeeeeeeeeegolang := new(Eeeeeeeeeegolang)

	err = json.Unmarshal(bytes, eeeeeeeeeegolang)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Eeeeeeeeeegolang")
	}

	return eeeeeeeeeegolang, nil
}

// UpdateEeeeeeeeeegolang retrieves an instance of Eeeeeeeeeegolang from the world state and updates its value
func (c *EeeeeeeeeegolangContract) UpdateEeeeeeeeeegolang(ctx contractapi.TransactionContextInterface, eeeeeeeeeegolangID string, newValue string) error {
	exists, err := c.EeeeeeeeeegolangExists(ctx, eeeeeeeeeegolangID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", eeeeeeeeeegolangID)
	}

	eeeeeeeeeegolang := new(Eeeeeeeeeegolang)
	eeeeeeeeeegolang.Value = newValue

	bytes, _ := json.Marshal(eeeeeeeeeegolang)

	return ctx.GetStub().PutState(eeeeeeeeeegolangID, bytes)
}

// DeleteEeeeeeeeeegolang deletes an instance of Eeeeeeeeeegolang from the world state
func (c *EeeeeeeeeegolangContract) DeleteEeeeeeeeeegolang(ctx contractapi.TransactionContextInterface, eeeeeeeeeegolangID string) error {
	exists, err := c.EeeeeeeeeegolangExists(ctx, eeeeeeeeeegolangID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", eeeeeeeeeegolangID)
	}

	return ctx.GetStub().DelState(eeeeeeeeeegolangID)
}
