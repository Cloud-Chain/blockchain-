/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"github.com/Cloud-Chain/blockchain-repo/inspection/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating vehicle inspection chaincodes: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting vehicle inspection chaincodes: %v", err)
	}
}
