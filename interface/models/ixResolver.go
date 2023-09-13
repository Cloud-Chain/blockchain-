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
func InitLedger(pc config.PeerConfig) {
	_, commit, err := pc.InspectionContract.SubmitAsync("InitLedger")
	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get InspectRequest transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit InspectRequest transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** InspectRequest committed successfully")
}
func InspectRequest(id int64, basicInfo BasicInfo, pc config.PeerConfig) Inspection {
	basicInfoJSON, err := json.Marshal(basicInfo)
	result, commit, err := pc.InspectionContract.SubmitAsync("InspectRequest",
		client.WithArguments(string(id), string(basicInfoJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit InspectRequest transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get InspectRequest transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit InspectRequest transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** InspectRequest committed successfully")
	fmt.Printf("resultJSON : %s \n", result)
	// result를 JSON 문자열로 변환
	//resultJSON, err := json.Marshal(result)
	//if err != nil {
	//	panic(fmt.Errorf("failed to marshal result to JSON: %s, %w", result, err))
	//}
	//fmt.Printf("resultJSON : %s", resultJSON)
	// resultJSON을 구조체로 언마샬링
	var resultStruct Inspection // YourResultStruct는 결과를 언마샬링할 구조체로 대체해야 합니다.
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func InspectResult(id int64, detailInfo DetailInfo, images Images, etc string, pc config.PeerConfig) Inspection {
	detailInfoJSON, err := json.Marshal(detailInfo)
	result, commit, err := pc.InspectionContract.SubmitAsync("InspectResult",
		client.WithArguments(string(id), string(detailInfoJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit InspectRequest transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get InspectRequest transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit InspectRequest transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** InspectRequest committed successfully")
	fmt.Printf("resultJSON : %s \n", result)
	// result를 JSON 문자열로 변환
	//resultJSON, err := json.Marshal(result)
	//if err != nil {
	//	panic(fmt.Errorf("failed to marshal result to JSON: %s, %w", result, err))
	//}
	//fmt.Printf("resultJSON : %s", resultJSON)
	// resultJSON을 구조체로 언마샬링
	var resultStruct Inspection // YourResultStruct는 결과를 언마샬링할 구조체로 대체해야 합니다.
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}
