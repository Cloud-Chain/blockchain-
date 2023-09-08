package models

type VehicleInfo struct {
	ID             int64           `json:"id"`
	InspectionDate string           `json:"inspectionDate"`
	VehicleBasicInfo   VehicleBasicInfo `json:"vehicleBasicInfo"`
	VehicleDetailInfo  VehicleDetailInfo `json:"vehicleDetailInfo"`
	Images             Images           `json:"images"`
	Etc                string           `json:"etc"`
}

type VehicleDetailInfo struct {
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

