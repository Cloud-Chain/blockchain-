package models

type CertRequest struct {
	Org      string `json:"org"`
	UserID   string `json:"userID"`
	Password string `json:"password"`
}
