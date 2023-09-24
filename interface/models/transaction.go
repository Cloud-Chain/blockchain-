package models

type Transaction struct {
	ID                 int64              `json:"id"`
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
