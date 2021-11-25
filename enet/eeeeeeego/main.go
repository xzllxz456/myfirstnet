/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	eeeeeeeeeegolangContract := new(EeeeeeeeeegolangContract)
	eeeeeeeeeegolangContract.Info.Version = "0.0.1"
	eeeeeeeeeegolangContract.Info.Description = "My Smart Contract"
	eeeeeeeeeegolangContract.Info.License = new(metadata.LicenseMetadata)
	eeeeeeeeeegolangContract.Info.License.Name = "Apache-2.0"
	eeeeeeeeeegolangContract.Info.Contact = new(metadata.ContactMetadata)
	eeeeeeeeeegolangContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(eeeeeeeeeegolangContract)
	chaincode.Info.Title = "eeeeeeego chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from EeeeeeeeeegolangContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
