/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"github.com/Cloud-Chain/blockchain-repo/transaction/chaincode"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating used car transfer chaincodes: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting used car transfer chaincodes: %v", err)
	}
}
