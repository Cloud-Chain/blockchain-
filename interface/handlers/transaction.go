package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interface/config"
	"interface/models"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
InitLedger
SellVehicle - transactionID int64, seller Participant, transactionDetails TransactionDetails
BuyVehicle -
QueryInspectionResult - inspectionID string
QueryAllInspectionRequest
*/
func SellVehicle(c *gin.Context) {
	var request models.Transaction
	org := c.GetHeader("org")
	userID := c.GetHeader("userID")
	cert := c.GetHeader("CA-User")
	if err := c.BindJSON(&request); err != nil {
		fmt.Println("err1")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if !checkCertForTransaction(org, userID, cert) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "잘못된 인증서입니다."})
		return
	}
	result := models.SellVehicle(request.ID, request.Assignor, request.TransactionDetails, config.SellerConfig)
	c.IndentedJSON(http.StatusCreated, result)
}
func BuyVehicle(c *gin.Context) {
	var request models.Transaction
	org := c.GetHeader("org")
	userID := c.GetHeader("userID")
	cert := c.GetHeader("CA-User")
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if !checkCertForTransaction(org, userID, cert) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "잘못된 인증서입니다."})
		return
	}
	result := models.BuyVehicle(request.ID, request.Assignee, request.TransactionDetails, config.BuyerConfig)
	c.IndentedJSON(http.StatusOK, result)
}
func SellerCompromiseTransaction(c *gin.Context) {
	var request models.Transaction
	org := c.GetHeader("org")
	userID := c.GetHeader("userID")
	cert := c.GetHeader("CA-User")
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if !checkCertForTransaction(org, userID, cert) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "잘못된 인증서입니다."})
		return
	}
	result := models.CompromiseTransaction(request.ID, request.TransactionDetails, config.SellerConfig, org)
	c.IndentedJSON(http.StatusOK, result)
}
func BuyerCompromiseTransaction(c *gin.Context) {
	var request models.Transaction
	org := c.GetHeader("org")
	userID := c.GetHeader("userID")
	cert := c.GetHeader("CA-User")
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if !checkCertForTransaction(org, userID, cert) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "잘못된 인증서입니다."})
		return
	}
	result := models.CompromiseTransaction(request.ID, request.TransactionDetails, config.BuyerConfig, org)
	c.IndentedJSON(http.StatusOK, result)
}
func ReadTransaction(c *gin.Context) {
	request := c.Query("id")
	result := models.ReadTransaction(request, config.InspectorConfig)
	c.IndentedJSON(http.StatusOK, result)
}
func QueryTransactionsByUser(c *gin.Context) {
	request := c.Query("userName")
	result := models.QueryTransactionsByUser(request, config.InspectorConfig)
	c.IndentedJSON(http.StatusOK, result)
}
func QueryTransactionsByVehicle(c *gin.Context) {
	request := c.Query("vehicleRegistrationNumber")
	result := models.QueryTransactionsByVehicle(request, config.InspectorConfig)
	c.IndentedJSON(http.StatusOK, result)
}
func QueryAllTransactions(c *gin.Context) {
	result := models.QueryAllTransactions(config.InspectorConfig)
	c.IndentedJSON(http.StatusOK, result)
}
func checkCertForTransaction(org, userID, cert string) bool {
	certPath := fmt.Sprintf("%s/%s.pnu.cse/users/%s@%s.pnu.cse/msp/signcerts/cert.pem", config.PROJECT_PATH, org, userID, org)

	// 파일 읽기
	fileContent, err := ioutil.ReadFile(certPath)

	// 모든 개행 문자를 제거
	convertedUserCert := strings.ReplaceAll(cert, "\"", "")
	convertedUserCert = strings.ReplaceAll(convertedUserCert, `\n`, "")
	convertedFileCert := strings.ReplaceAll(string(fileContent), "\n", "")
	fmt.Println("request cert : " + convertedUserCert)
	fmt.Println("saved cert in blockchain : " + convertedFileCert)
	if err != nil {
		return false
	} else if convertedFileCert != convertedUserCert {
		return false
	}
	return true
}
