package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/common/flogging"
)

type SmartContract struct {
	contractapi.Contract
}

var logger = flogging.MustGetLogger("mychaincode")

type Inspection struct {
	ID                int64      `json:"id"`
	RequestDate       string     `json:"requestDate"`
	InspectionDate    string     `json:"inspectionDate"`
	VehicleBasicInfo  BasicInfo  `json:"vehicleBasicInfo"`
	VehicleDetailInfo DetailInfo `json:"vehicleDetailInfo"`
	Images            Images     `json:"images"`
	Etc               string     `json:"etc"`
}

type BasicInfo struct {
	VehicleIdentificationNumber string `json:"vehicleIdentificationNumber"`
	VehicleModelName            string `json:"vehicleModelName"`
	VehicleRegistrationNumber   string `json:"vehicleRegistrationNumber"`
	GearboxType                 string `json:"gearboxType"`
	FuelUsed                    string `json:"fuelUsed"`
	Mileage                     int    `json:"mileage"`
	Color                       string `json:"color"`
	Options                     string `json:"options"`
}

type DetailInfo struct {
	Tuning       string `json:"tuning"`
	OuterPlate   string `json:"outerPlate"`
	VehicleFrame string `json:"vehicleFrame"`
	Motor        string `json:"motor"`
	Transmission string `json:"transmission"`
	Steering     string `json:"steering"`
	Braking      string `json:"braking"`
	Electricity  string `json:"electricity"`
	Fuel         string `json:"fuel"`
	Exterior     string `json:"exterior"`
	Interior     string `json:"interior"`
	Gloss        string `json:"gloss"`
	Wheel        string `json:"wheel"`
	Tire         string `json:"tire"`
	Glass        string `json:"glass"`
}
type Images struct {
	Inside  string `json:"inside"`
	Outside string `json:"outside"`
	Front   string `json:"front"`
	Left    string `json:"left"`
	Right   string `json:"right"`
	Back    string `json:"back"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	inspection := Inspection{
		ID:             0,
		RequestDate:    "2023-07-01",
		InspectionDate: "2023-07-26",
		VehicleBasicInfo: BasicInfo{
			VehicleIdentificationNumber: "1HGCM82633A123456",
			VehicleModelName:            "Toyota Camry",
			VehicleRegistrationNumber:   "ABCD123",
			GearboxType:                 "Automatic",
			FuelUsed:                    "Gasoline",
			Mileage:                     10000,
			Color:                       "Silver",
			Options:                     "Navigation, Leather Seats, Sunroof",
		},
		VehicleDetailInfo: DetailInfo{
			Tuning:       "None",
			OuterPlate:   "Good",
			VehicleFrame: "Intact",
			Motor:        "Engine in good condition",
			Transmission: "Smooth",
			Steering:     "Responsive",
			Braking:      "Effective",
			Electricity:  "All electrical systems functional",
			Fuel:         "No leaks or issues",
			Exterior:     "Clean and well-maintained",
			Interior:     "Neat and tidy",
			Gloss:        "Shiny",
			Wheel:        "Good condition",
			Tire:         "Adequate tread depth",
			Glass:        "No cracks or chips",
		},
		Images: Images{
			Inside:  "https://example.com/car_images/inside.jpg",
			Outside: "https://example.com/car_images/outside.jpg",
			Front:   "https://example.com/car_images/front.jpg",
			Left:    "https://example.com/car_images/left.jpg",
			Right:   "https://example.com/car_images/right.jpg",
			Back:    "https://example.com/car_images/back.jpg",
		},
		Etc: "I think...",
	}

	inspectionJSON, err := json.Marshal(inspection)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(strconv.FormatInt(inspection.ID, 10), inspectionJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state: %v", err)
	}
	err = ctx.GetStub().PutState("lastInspectionID", []byte(strconv.FormatInt(inspection.ID, 10)))
	if err != nil {
		return fmt.Errorf("Failed to update lastInspectionID in world state: %s", err.Error())
	}

	return nil
}
func (s *SmartContract) InspectRequest(ctx contractapi.TransactionContextInterface, basicInfo BasicInfo) (*Inspection, error) {
	lastInspectionID, err := ctx.GetStub().GetState("lastInspectionID")
	lastID := int64(0)
	if lastInspectionID != nil {
		lastID, _ = strconv.ParseInt(string(lastInspectionID), 10, 64)
	} else {
		lastID = 0
	}
	inspection := Inspection{
		ID:               lastID + 1,
		RequestDate:      time.Now().Format("2006-01-02 15:04:05"),
		InspectionDate:   "",
		VehicleBasicInfo: basicInfo,
		VehicleDetailInfo: DetailInfo{
			Tuning:       "",
			OuterPlate:   "",
			VehicleFrame: "",
			Motor:        "",
			Transmission: "",
			Steering:     "",
			Braking:      "",
			Electricity:  "",
			Fuel:         "",
			Exterior:     "",
			Interior:     "",
			Gloss:        "",
			Wheel:        "",
			Tire:         "",
			Glass:        "",
		},
		Images: Images{
			Inside:  "",
			Outside: "",
			Front:   "",
			Left:    "",
			Right:   "",
			Back:    "",
		},
		Etc: "",
	}

	inspectionAsBytes, err := json.Marshal(inspection)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshaling transaction as bytes. %s", err.Error())
	}

	err = ctx.GetStub().PutState(strconv.FormatInt(inspection.ID, 10), inspectionAsBytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to put transaction to world state. %s", err.Error())
	}
	if err != nil {
		log.Println(err.Error())
	}
	err = ctx.GetStub().PutState("lastInspectionID", []byte(strconv.FormatInt(inspection.ID, 10)))
	if err != nil {
		return nil, fmt.Errorf("Failed to update lastInspectionID in world state: %s", err.Error())
	}

	return &inspection, nil
}

func (s *SmartContract) InspectResult(
	ctx contractapi.TransactionContextInterface, inspectionID string, detailInfo DetailInfo, images Images, etc string) (*Inspection, error) {

	originInspectionData, err := ctx.GetStub().GetState(inspectionID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if originInspectionData == nil {
		return nil, fmt.Errorf("%s does not exist", inspectionID)
	}

	inspection := new(Inspection)
	_ = json.Unmarshal(originInspectionData, inspection)

	if inspection.InspectionDate != "" {
		return nil, fmt.Errorf("%s is already inspected", inspectionID)
	}
	inspection.InspectionDate = time.Now().Format("2006-01-02 15:04:05")
	inspection.Images = images
	inspection.VehicleDetailInfo = detailInfo
	inspection.Etc = etc

	newInspectionData, _ := json.Marshal(inspection)
	err = ctx.GetStub().PutState(strconv.FormatInt(inspection.ID, 10), newInspectionData)
	if err != nil {
		return nil, fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	return inspection, nil
}

func (s *SmartContract) QueryInspectionResult(ctx contractapi.TransactionContextInterface, inspectionID string) (*Inspection, error) {
	inspectionData, err := ctx.GetStub().GetState(inspectionID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if inspectionData == nil {
		return nil, fmt.Errorf("%s does not exist", inspectionID)
	}

	inspection := new(Inspection)
	_ = json.Unmarshal(inspectionData, inspection)

	return inspection, nil
}

func (s *SmartContract) QueryAllInspectionRequest(ctx contractapi.TransactionContextInterface) ([]*Inspection, error) {
	queryString := fmt.Sprintf(`{
		"selector": {
			"_id": {
				"$gt": null
			}
		}
	}`)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		logger.Errorf("Error querying data: %v", err)
		return nil, err
	}
	defer resultsIterator.Close()

	var inspections []*Inspection
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Errorf("Error reading query response: %v", err)
			return nil, err
		}
		var item Inspection
		if err = json.Unmarshal(queryResponse.Value, &item); err != nil {
			logger.Errorf("Error unmarshaling data: %v", err)
			return nil, err
		}

		inspections = append(inspections, &item)
	}

	// 로그에 데이터 기록
	logger.Infof("Retrieved %d inspection records", len(inspections))

	return inspections, nil
}
