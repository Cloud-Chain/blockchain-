package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Inspection struct {
	ID                int        `json:"id"`
	InspectionStatus  bool       `json:"inspectionStatus"`
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
		ID:               0,
		InspectionStatus: false,
		RequestDate:      "2023-07-01",
		InspectionDate:   "2023-07-26",
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

	err = ctx.GetStub().PutState(strconv.Itoa(inspection.ID), inspectionJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state: %v", err)
	}
	err = ctx.GetStub().PutState("lastInspectionID", []byte(strconv.Itoa(inspection.ID)))
	if err != nil {
		return fmt.Errorf("Failed to update lastInspectionID in world state: %s", err.Error())
	}

	return nil
}
func (s *SmartContract) InspectRequest(ctx contractapi.TransactionContextInterface, basicInfo BasicInfo) (*Inspection, error) {
	creator,err := ctx.GetStub().GetCreator()
	if err != nil {
		return nil, fmt.Errorf("Failed to get creator : %v", err)
	}
	sId := &msp.SerializedIdentity{}
	err = proto.Unmarshal(creator,sId)
	if sId.GetMspid() != "sellerMSP" {
		return nil, fmt.Errorf("%v does not have permission for vehicle inspection requests.\n", sId.GetMspid())
	}

	lastInspectionID, err := ctx.GetStub().GetState("lastInspectionID")
	lastID := int(0)
	if lastInspectionID != nil {
		lastID, _ = strconv.Atoi(string(lastInspectionID))
	} else {
		lastID = 0
	}
	inspection := Inspection{
		ID:               lastID+1,
		InspectionStatus: false,
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

	err = ctx.GetStub().PutState(strconv.Itoa(inspection.ID), inspectionAsBytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to put transaction to world state. %s", err.Error())
	}
	if err != nil {
		log.Println(err.Error())
	}
	err = ctx.GetStub().PutState("lastInspectionID", []byte(strconv.Itoa(inspection.ID)))
	if err != nil {
		return nil, fmt.Errorf("Failed to update lastInspectionID in world state: %s", err.Error())
	}

	return &inspection, nil
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

func (s *SmartContract) InspectResult(ctx contractapi.TransactionContextInterface, inspection Inspection) (*Inspection, error) {
	creator,err := ctx.GetStub().GetCreator()
	if err != nil {
		return nil, fmt.Errorf("Failed to get creator : %v", err)
	}
	sId := &msp.SerializedIdentity{}
	err = proto.Unmarshal(creator,sId)
	if sId.GetMspid() != "inspectorMSP" {
		return nil, fmt.Errorf("%v does not have permission for update vehicle inspection results.\n", sId.GetMspid())
	}
	
	inspection.InspectionStatus = true
	inspection.RequestDate = time.Now().Format("2006-01-02 15:04:05")

	inspectionAsBytes, err := json.Marshal(inspection)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshaling transaction as bytes. %s", err.Error())
	}
	err = ctx.GetStub().PutState(strconv.Itoa(inspection.ID), inspectionAsBytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to put transaction to world state. %s", err.Error())
	}
	if err != nil {
		log.Println(err.Error())
	}
	return &inspection, nil
}

// func (s *SmartContract) InspectResult(
// 	ctx contractapi.TransactionContextInterface, inspectionID string, detailInfo DetailInfo, images Images, etc string) (*Inspection, error) {

// 	originInspectionData, err := ctx.GetStub().GetState(inspectionID)

// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
// 	}

// 	if originInspectionData == nil {
// 		return nil, fmt.Errorf("%s does not exist", inspectionID)
// 	}

// 	inspection := new(Inspection)
// 	err = json.Unmarshal(originInspectionData, inspection)
// 	if err != nil {
// 		fmt.Errorf("+v\n", err)
// 	}

// 	if inspection.InspectionDate != "" {
// 		return nil, fmt.Errorf("%s is already inspected", inspectionID)
// 	}
// 	inspection.InspectionStatus = true
// 	inspection.InspectionDate = time.Now().Format("2006-01-02 15:04:05")
// 	inspection.Images = images
// 	inspection.VehicleDetailInfo = detailInfo
// 	inspection.Etc = etc

// 	newInspectionData, _ := json.Marshal(inspection)
// 	err = ctx.GetStub().PutState(strconv.Itoa(inspection.ID), newInspectionData)
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to put to world state. %s", err.Error())
// 	}

// 	return inspection, nil
// }

func (s *SmartContract) QueryAllInspections(ctx contractapi.TransactionContextInterface) ([]Inspection, error) {
	queryString := `{
		"selector": {
		   "vehicleBasicInfo.mileage": {
			  "$gte": 0
		   }
		}
	 }`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("Failed to get query result: %s", err.Error())
	}
	defer resultsIterator.Close()

	inspections := []Inspection{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var inspection Inspection

		err = json.Unmarshal(queryResponse.Value, &inspection)
		if err != nil {
			fmt.Printf("Unmarshal Error : %+v , %+v \n", queryResponse.Value, err)
			// return nil, err
		}

		inspections = append(inspections, inspection)
		fmt.Printf("%+v \n", queryResponse.Value)
		fmt.Printf("%+v \n", inspection)
	}
	return inspections, nil
}
