package handlers

import (
	"github.com/gin-gonic/gin"
	"interface/config"
	"interface/models"
	"net/http"
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
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.SellVehicle(request.ID, request.Assignor, request.TransactionDetails, config.SellerConfig)
	c.IndentedJSON(http.StatusCreated, result)
}
func BuyVehicle(c *gin.Context) {
	var request models.Transaction
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.BuyVehicle(request.ID, request.Assignee, request.TransactionDetails, config.BuyerConfig)
	c.IndentedJSON(http.StatusOK, result)
}
func SellerCompromiseTransaction(c *gin.Context) {
	var request models.Transaction
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.CompromiseTransaction(request.ID, request.TransactionDetails, config.SellerConfig)
	c.IndentedJSON(http.StatusOK, result)
}
func BuyerCompromiseTransaction(c *gin.Context) {
	var request models.Transaction
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.CompromiseTransaction(request.ID, request.TransactionDetails, config.BuyerConfig)
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
