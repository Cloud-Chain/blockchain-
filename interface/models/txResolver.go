package models

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"interface/config"
	"strconv"
)

/*
InitLedger
SellVehicle - transactionID string, seller Participant, transactionDetails TransactionDetails
BuyVehicle -
QueryInspectionResult - inspectionID string
QueryAllInspectionRequest
*/
func TransactionInitLedger(pc config.PeerConfig) {
	_, commit, err := pc.TransactionContract.SubmitAsync("InitLedger")
	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get Transaction InitLedger transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit Transaction InitLedger transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** Transaction InitLedger committed successfully")
}
func SellVehicle(id int64, seller Participant, details TransactionDetails, pc config.PeerConfig) Transaction {
	sellerJSON, err := json.Marshal(seller)
	detailsJSON, err := json.Marshal(details)

	result, commit, err := pc.TransactionContract.SubmitAsync("SellVehicle",
		client.WithArguments(strconv.FormatInt(id, 10), string(sellerJSON), string(detailsJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit SellVehicle transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get SellVehicle transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit SellVehicle transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** SellVehicle committed successfully")
	printStatus(status)
	

	var resultStruct Transaction // YourResultStruct는 결과를 언마샬링할 구조체로 대체
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	// resultJSON, _ := json.MarshalIndent(resultStruct, "", "    ")
	// fmt.Printf("resultStruct:\n%s\n", resultJSON)
	return resultStruct
}

func BuyVehicle(id int64, buyer Participant, details TransactionDetails, pc config.PeerConfig) Transaction {
	buyerJSON, err := json.Marshal(buyer)
	detailsJSON, err := json.Marshal(details)

	result, commit, err := pc.TransactionContract.SubmitAsync("BuyVehicle",
		client.WithArguments(strconv.FormatInt(id, 10), string(buyerJSON), string(detailsJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit BuyVehicle transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get BuyVehicle transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit BuyVehicle transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** BuyVehicle committed successfully")
	printStatus(status)

	var resultStruct Transaction // YourResultStruct는 결과를 언마샬링할 구조체로 대체
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	// fmt.Printf("resultStruct %+v", resultStruct)
	// resultJSON, _ := json.MarshalIndent(resultStruct, "", "    ")
	// fmt.Printf("resultStruct:\n%s\n", resultJSON)
	return resultStruct
}
func CompromiseTransaction(id int64, details TransactionDetails, pc config.PeerConfig ,org string) Transaction {
	if details.TransactionState != "Accept" && org == "seller" {
		details.TransactionState = "SellerRequest"
	} 
	if details.TransactionState != "Accept" && org == "buyer" {
		details.TransactionState = "BuyerRequest"
	} 
	
	detailsJSON, err := json.Marshal(details)
	
	result, commit, err := pc.TransactionContract.SubmitAsync("CompromiseTransaction",
		client.WithArguments(strconv.FormatInt(id, 10), string(detailsJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit CompromiseTransaction transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get CompromiseTransaction transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit CompromiseTransaction transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** CompromiseTransaction committed successfully")
	printStatus(status)

	var resultStruct Transaction // YourResultStruct는 결과를 언마샬링할 구조체로 대체
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	// fmt.Printf("resultStruct %+v", resultStruct)
	// resultJSON, _ := json.MarshalIndent(resultStruct, "", "    ")
	// fmt.Printf("resultStruct:\n%s\n", resultJSON)
	return resultStruct
}
func ReadTransaction(id string, pc config.PeerConfig) Transaction {
	result, err := pc.TransactionContract.EvaluateTransaction("ReadTransaction", id)
	if err != nil {
		panic(fmt.Errorf("failed to query ReadTransaction: %w", err))
	}

	fmt.Println("\n*** ReadTransaction successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct Transaction
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}

	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}
func QueryTransactionsByUser(userName string, pc config.PeerConfig) []Transaction {
	result, err := pc.TransactionContract.EvaluateTransaction("QueryTransactionsByUser", userName)
	if err != nil {
		panic(fmt.Errorf("failed to query QueryTransactionsByUser: %w", err))
	}

	fmt.Println("\n*** QueryTransactionsByUser successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct []Transaction
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}

	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}
func QueryTransactionsByVehicle(vehicleRegistrationNumber string, pc config.PeerConfig) []Transaction {
	result, err := pc.TransactionContract.EvaluateTransaction("QueryTransactionsByVehicle", vehicleRegistrationNumber)
	if err != nil {
		panic(fmt.Errorf("failed to query QueryTransactionsByVehicle: %w", err))
	}

	fmt.Println("\n*** QueryTransactionsByVehicle successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct []Transaction
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}

	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}
func QueryAllTransactions(pc config.PeerConfig) []Transaction {
	result, err := pc.TransactionContract.EvaluateTransaction("QueryAllTransactions")
	if err != nil {
		panic(fmt.Errorf("failed to query QueryAllTransactions: %w", err))
	}

	fmt.Println("\n*** QueryAllTransactions successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct []Transaction
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}
	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}
