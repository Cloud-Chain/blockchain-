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

type Transaction struct {
	ID                 int              `json:"id"`
	UploadDate         string             `json:"uploadDate"`
	Assignor           Participant        `json:"assignor"`
	Assignee           Participant        `json:"assignee"`
	TransactionDetails TransactionDetails `json:"transactionDetails"`
}

type Participant struct {
	Name                       string `json:"name"`
	ResidentRegistrationNumber string `json:"residentRegistrationNumber"`
	PhoneNumber                string `json:"phoneNumber"`
	Address                    string `json:"address"`
}

type TransactionDetails struct {
	TransactionState             string `json:"transactionState"`
	VehicleRegistrationNumber    string `json:"vehicleRegistrationNumber"`
	NewVehicleRegistrationNumber string `json:"newVehicleRegistrationNumber"`
	VehicleModelName             string `json:"vehicleModelName"`
	VehicleIdentificationNumber  string `json:"vehicleIdentificationNumber"`
	TransactionDate              string `json:"transactionDate"`
	TransactionAmount            int `json:"transactionAmount"`
	BalancePaymentDate           string `json:"balancePaymentDate"`
	VehicleDeliveryDate          string `json:"vehicleDeliveryDate"`
	VehicleDeliveryAddress       string `json:"vehicleDeliveryAddress"`
	Mileage                      int `json:"mileage"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	transactions := []Transaction{
		{
			ID:         999999999,
			UploadDate: "2023-06-23 00:00:00",
			Assignor:   Participant{Name: "test1", ResidentRegistrationNumber: "1111111-1111111", PhoneNumber: "111-1111-1111", Address: "1111 TEST St"},
			Assignee:   Participant{},
			TransactionDetails: TransactionDetails{
				TransactionState:             "SellerRequest",
				VehicleRegistrationNumber:    "111A1111",
				NewVehicleRegistrationNumber: "",
				VehicleModelName:             "Tesla Model A - 1111",
				VehicleIdentificationNumber:  "5YJ3E1EA1JF11111",
				TransactionDate:              "",
				TransactionAmount:            10000,
				BalancePaymentDate:           "",
				VehicleDeliveryDate:          "",
				VehicleDeliveryAddress:       "",
				Mileage:                      10000,
			},
		},
		{
			ID:         999999998,
			UploadDate: "2023-06-23 00:00:00",
			Assignor:   Participant{Name: "test2", ResidentRegistrationNumber: "222222-2222222", PhoneNumber: "222-2222-2222", Address: "2222 TEST St"},
			Assignee:   Participant{},
			TransactionDetails: TransactionDetails{
				TransactionState:             "SellerRequest",
				VehicleRegistrationNumber:    "222B2222",
				NewVehicleRegistrationNumber: "",
				VehicleModelName:             "Tesla Model S - 2222",
				VehicleIdentificationNumber:  "5YJ3E1EA1JF22222",
				TransactionDate:              "",
				TransactionAmount:            10000,
				BalancePaymentDate:           "",
				VehicleDeliveryDate:          "",
				VehicleDeliveryAddress:       "",
				Mileage:                      20000,
			},
		},
		{
			ID:         999999997,
			UploadDate: "2023-06-23 00:00:00",
			Assignor:   Participant{Name: "test3", ResidentRegistrationNumber: "333333-3333333", PhoneNumber: "333-3333-3333", Address: "3333 TEST St"},
			Assignee:   Participant{Name: "test4", ResidentRegistrationNumber: "444444-4444444", PhoneNumber: "444-4444-4444", Address: "4444 TEST St"},
			TransactionDetails: TransactionDetails{
				TransactionState:             "SoldOut",
				VehicleRegistrationNumber:    "333C3333",
				NewVehicleRegistrationNumber: "444C4444",
				VehicleModelName:             "Tesla Model S - 3333",
				VehicleIdentificationNumber:  "5YJ3E1EA1JF33333",
				TransactionDate:              "2023-06-25 00:00:00",
				TransactionAmount:            1000,
				BalancePaymentDate:           "2023-06-27 00:00:00",
				VehicleDeliveryDate:          "2023-06-27 00:00:00",
				VehicleDeliveryAddress:       "4444 TEST St",
				Mileage:                      30000,
			},
		},
		{
			ID:         999999996,
			UploadDate: "2023-06-23 00:00:00",
			Assignor:   Participant{Name: "test3", ResidentRegistrationNumber: "333333-3333333", PhoneNumber: "333-3333-3333", Address: "3333 TEST St"},
			Assignee:   Participant{Name: "test4", ResidentRegistrationNumber: "444444-4444444", PhoneNumber: "444-4444-4444", Address: "4444 TEST St"},
			TransactionDetails: TransactionDetails{
				TransactionState:             "BuyerRequest",
				VehicleRegistrationNumber:    "333C3333",
				NewVehicleRegistrationNumber: "444C4444",
				VehicleModelName:             "Tesla Model S - 3333",
				VehicleIdentificationNumber:  "5YJ3E1EA1JF33333",
				TransactionDate:              "2023-06-25 00:00:00",
				TransactionAmount:            1000,
				BalancePaymentDate:           "2023-06-27 00:00:00",
				VehicleDeliveryDate:          "2023-06-27 00:00:00",
				VehicleDeliveryAddress:       "4444 TEST St",
				Mileage:                      30000,
			},
		},
		// Add more transactions if required...
	}

	for _, transaction := range transactions {
		transactionData, _ := json.Marshal(transaction)
		err := ctx.GetStub().PutState(strconv.Itoa(transaction.ID), transactionData)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}
func (s *SmartContract) SellVehicle(ctx contractapi.TransactionContextInterface, transactionID string, seller Participant, transactionDetails TransactionDetails) (*Transaction, error) {
	creator,err := ctx.GetStub().GetCreator()
	if err != nil {
		return nil, fmt.Errorf("Failed to get creator : %v", err)
	}
	sId := &msp.SerializedIdentity{}
	err = proto.Unmarshal(creator,sId)
	if sId.GetMspid() != "sellerMSP" {
		return nil, fmt.Errorf("%v does not have permission for update vehicle inspection results.\n", sId.GetMspid())
	}
	
	lastTransactionID, err := ctx.GetStub().GetState("lastTransactionID")
	lastID := int(0)
	if lastTransactionID != nil {
		lastID, _ = strconv.Atoi(string(lastTransactionID))
	} else {
		lastID = 0
	}
	// Assign new ID to the transaction
	newTransactionDetails := TransactionDetails{
		TransactionState:             "SellerRequest",
		VehicleRegistrationNumber:    transactionDetails.VehicleRegistrationNumber,
		NewVehicleRegistrationNumber: "",
		VehicleModelName:             transactionDetails.VehicleModelName,
		VehicleIdentificationNumber:  transactionDetails.VehicleIdentificationNumber,
		TransactionDate:              "",
		TransactionAmount:            transactionDetails.TransactionAmount,
		BalancePaymentDate:           "",
		VehicleDeliveryDate:          "",
		VehicleDeliveryAddress:       "",
		Mileage:                      transactionDetails.Mileage,
	}
	// parseID, err := strconv.ParseInt(transactionID, 10, 64)
	// if err != nil {
	// 	fmt.Errorf("Failed to parsing transactionID as int64. %s", err.Error())
	// }
	transaction := Transaction{
		ID:                 lastID+1,
		UploadDate:         time.Now().Format("2006-01-02 15:04:05"),
		Assignor:           seller,
		Assignee:           Participant{},
		TransactionDetails: newTransactionDetails,
	}

	transactionAsBytes, err := json.Marshal(transaction)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshaling transaction as bytes. %s", err.Error())
	}

	err = ctx.GetStub().PutState(strconv.Itoa(transaction.ID), transactionAsBytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to put transaction to world state. %s", err.Error())
	}
	err = ctx.GetStub().SetEvent("SellVehicle", transactionAsBytes)
	if err != nil {
		log.Println(err.Error())
	}
	err = ctx.GetStub().PutState("lastTransactionID", []byte(strconv.Itoa(transaction.ID)))
	if err != nil {
		return nil, fmt.Errorf("Failed to update lastTransactionID in world state: %s", err.Error())
	}

	return &transaction, nil
}

func (s *SmartContract) BuyVehicle(ctx contractapi.TransactionContextInterface, transactionID string, buyer Participant, transactionDetails TransactionDetails) (*Transaction, error) {
	creator,err := ctx.GetStub().GetCreator()
	if err != nil {
		return nil, fmt.Errorf("Failed to get creator : %v", err)
	}
	sId := &msp.SerializedIdentity{}
	err = proto.Unmarshal(creator,sId)
	if sId.GetMspid() != "buyerMSP" {
		return nil, fmt.Errorf("%v does not have permission for update vehicle inspection results.\n", sId.GetMspid())
	}
		
	oldTransactionData, err := ctx.GetStub().GetState(transactionID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if oldTransactionData == nil {
		return nil, fmt.Errorf("%s does not exist", transactionID)
	}

	transaction := new(Transaction)
	_ = json.Unmarshal(oldTransactionData, transaction)

	if transaction.TransactionDetails.TransactionState == "SoldOut" {
		return nil, fmt.Errorf("Vehicle is already sold")
	}
	transaction.Assignee = buyer
	transaction.TransactionDetails.TransactionState = "BuyerRequest"
	transaction.TransactionDetails.NewVehicleRegistrationNumber = transactionDetails.NewVehicleRegistrationNumber
	// transaction.TransactionDetails.TradingDate =
	transaction.TransactionDetails.BalancePaymentDate = transactionDetails.BalancePaymentDate
	transaction.TransactionDetails.VehicleDeliveryDate = transactionDetails.VehicleDeliveryDate
	transaction.TransactionDetails.VehicleDeliveryAddress = transactionDetails.VehicleDeliveryAddress

	newTransactionData, _ := json.Marshal(transaction)
	err = ctx.GetStub().PutState(strconv.Itoa(transaction.ID), newTransactionData)
	if err != nil {
		return nil, fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	err = ctx.GetStub().SetEvent("BuyVehicle", newTransactionData)
	if err != nil {
		log.Println(err.Error())
	}

	return transaction, nil
}

func (s *SmartContract) ReadTransaction(ctx contractapi.TransactionContextInterface, transactionID string) (*Transaction, error) {
	
	transactionData, err := ctx.GetStub().GetState(transactionID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if transactionData == nil {
		return nil, fmt.Errorf("%s does not exist", transactionID)
	}

	transaction := new(Transaction)
	_ = json.Unmarshal(transactionData, transaction)

	return transaction, nil
}

func (s *SmartContract) CompromiseTransaction(
	ctx contractapi.TransactionContextInterface, transactionID string, transactionDetails TransactionDetails) (*Transaction, error) {
	oldTransactionData, err := ctx.GetStub().GetState(transactionID)

	creator,err := ctx.GetStub().GetCreator()
	if err != nil {
		return nil, fmt.Errorf("Failed to get creator : %v", err)
	}
	sId := &msp.SerializedIdentity{}
	err = proto.Unmarshal(creator,sId)
	if sId.GetMspid() == "inspectorMSP" {
		return nil, fmt.Errorf("%v does not have permission for update vehicle inspection results.\n", sId.GetMspid())
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if oldTransactionData == nil {
		return nil, fmt.Errorf("%s does not exist", transactionID)
	}

	transaction := new(Transaction)
	_ = json.Unmarshal(oldTransactionData, transaction)

	if transaction.TransactionDetails.TransactionState == "SoldOut" {
		return nil, fmt.Errorf("Vehicle is already sold")
	} else if transactionDetails.TransactionState == "Accept" {
		transaction.TransactionDetails.TransactionState = "SoldOut"
		transaction.TransactionDetails.TransactionDate = time.Now().Format("2006-01-02 15:04:05")

	} else {
		transaction.TransactionDetails.TransactionState = transactionDetails.TransactionState
		transaction.TransactionDetails.NewVehicleRegistrationNumber = transactionDetails.NewVehicleRegistrationNumber
		transaction.TransactionDetails.VehicleDeliveryAddress = transactionDetails.VehicleDeliveryAddress
		transaction.TransactionDetails.VehicleDeliveryDate = transactionDetails.VehicleDeliveryDate
		transaction.TransactionDetails.BalancePaymentDate = transactionDetails.BalancePaymentDate
	}

	newTransactionData, _ := json.Marshal(transaction)
	err = ctx.GetStub().PutState(strconv.Itoa(transaction.ID), newTransactionData)
	if err != nil {
		return nil, fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	err = ctx.GetStub().SetEvent("CompromiseTransaction", newTransactionData)
	if err != nil {
		log.Println(err.Error())
	}
	err = ctx.GetStub().PutState(strconv.Itoa(transaction.ID), newTransactionData)
	if err != nil {
		return nil, fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	return transaction, nil
}

func (s *SmartContract) QueryTransactionsByUser(ctx contractapi.TransactionContextInterface, userName string) ([]*Transaction, error) {
	queryString := fmt.Sprintf(`{"selector":{"$or":[{"assignor.name":"%s"},{"assignee.name":"%s"}]}}`, userName, userName)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("Failed to get query result: %s", err.Error())
	}
	defer resultsIterator.Close()

	transactions := []*Transaction{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var transaction Transaction
		err = json.Unmarshal(queryResponse.Value, &transaction)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

func (s *SmartContract) QueryTransactionsByVehicle(ctx contractapi.TransactionContextInterface, vehicleRegistrationNumber string) ([]*Transaction, error) {
	queryString := fmt.Sprintf(`{"selector":{"transactionDetails.vehicleRegistrationNumber":"%s"}}`, vehicleRegistrationNumber)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("Failed to get query result: %s", err.Error())
	}
	defer resultsIterator.Close()

	transactions := []*Transaction{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var transaction Transaction
		err = json.Unmarshal(queryResponse.Value, &transaction)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

func (s *SmartContract) QueryAllTransactions(ctx contractapi.TransactionContextInterface) ([]Transaction, error) {
	// Empty selector for CouchDB - this will get all Transactions from the ledger
	queryString := `{
		"selector": {
		   "transactionDetails.transactionAmount": {
			  "$gte": 0
		   }
		}
	 }`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []Transaction
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var transaction Transaction
		if err = json.Unmarshal(queryResponse.Value, &transaction); err != nil {
			return nil, err
		}
		fmt.Printf("%v\n", transaction)
		fmt.Printf("%v\n", transactions)
		transactions = append(transactions, transaction)
		fmt.Printf("%v\n", transactions)
	}

	return transactions, nil
}
