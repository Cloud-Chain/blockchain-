package models

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"interface/config"
)

/*
InitLedger
InspectRequest - id int64, basicInfo BasicInfo
InspectResult - inspectionID string, detailInfo DetailInfo, images Images, etc string
QueryInspectionResult - inspectionID string
QueryAllInspectionRequest
*/
func InspectionInitLedger(pc config.PeerConfig) {
	_, commit, err := pc.InspectionContract.SubmitAsync("InitLedger")
	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get Inspection InitLedger transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit Inspection InitLedger transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** Inspection InitLedger committed successfully")
}
func InspectRequest(basicInfo BasicInfo, pc config.PeerConfig) Inspection {
	basicInfoJSON, err := json.Marshal(basicInfo)
	result, commit, err := pc.InspectionContract.SubmitAsync("InspectRequest",
		client.WithArguments(string(basicInfoJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit InspectRequest transaction: %+v", err))	
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get InspectRequest transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit InspectRequest transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** InspectRequest committed successfully")
	printStatus(status)

	var resultStruct Inspection 
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	// fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func InspectResult(inspection Inspection, pc config.PeerConfig) Inspection {
	inspectionJSON, err := json.Marshal(inspection)
	result, commit, err := pc.InspectionContract.SubmitAsync("InspectResult",
		client.WithArguments(string(inspectionJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit InspectResult transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get InspectResult transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit InspectResult transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** InspectResult committed successfully")
	printStatus(status)

	var resultStruct Inspection // YourResultStruct는 결과를 언마샬링할 구조체로 대체
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	// fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func QueryInspectResult(id string, pc config.PeerConfig) Inspection {
	result, err := pc.InspectionContract.EvaluateTransaction("QueryInspectionResult", id)
	if err != nil {
		panic(fmt.Errorf("failed to query QueryInspectionResult: %w", err))
	}

	fmt.Println("\n*** QueryInspectResult successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct Inspection
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}

	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func QueryAllInspections(pc config.PeerConfig) []Inspection {
	result, err := pc.InspectionContract.EvaluateTransaction("QueryAllInspections")
	if err != nil {
		panic(fmt.Errorf("failed to query QueryAllInspections: %w", err))
	}

	fmt.Println("\n*** QueryAllInspections successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct []Inspection
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}

	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func printStatus(s *client.Status) {
	fmt.Printf("Status:\n")
	fmt.Printf("  Code: %v\n", s.Code)
	fmt.Printf("  Successful: %v\n", s.Successful)
	fmt.Printf("  TransactionID: %v\n", s.TransactionID)
	fmt.Printf("  BlockNumber: %v\n", s.BlockNumber)
	fmt.Println()
}
