package models

type Inspection struct {
	ID               int      `json:"id"`
	InspectionStatus bool       `json:"inspectionStatus"`
	RequestDate      string     `json:"requestDate"`
	InspectionDate   string     `json:"inspectionDate"`
	BasicInfo        BasicInfo  `json:"vehicleBasicInfo"`
	DetailInfo       DetailInfo `json:"vehicleDetailInfo"`
	Images           Images     `json:"images"`
	Etc              string     `json:"etc"`
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
