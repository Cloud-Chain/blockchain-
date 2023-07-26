/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating used car transfer chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting used car transfer chaincode: %v", err)
	}
}
