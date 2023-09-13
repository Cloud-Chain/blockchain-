package handlers

import (
	"github.com/gin-gonic/gin"
	"interface/config"
	"interface/models"
	"net/http"
	"strconv"
)

/*
InitLedger
InspectRequest - id int64, basicInfo BasicInfo
InspectResult - inspectionID string, detailInfo DetailInfo, images Images, etc string
QueryInspectionResult - inspectionID string
QueryAllInspectionRequest
*/
func RequestInspection(c *gin.Context) {
	var request models.Inspection
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.InspectRequest(request.BasicInfo, config.SellerConfig)
	c.IndentedJSON(http.StatusCreated, result)
}

func ExecuteInspection(c *gin.Context) {
	var request models.Inspection
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.InspectResult(request.ID, request.DetailInfo, request.Images, request.Etc, config.SellerConfig)
	c.IndentedJSON(http.StatusOK, result)
}

func FindInspection(c *gin.Context) {
	request, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
	}
	result := models.QueryInspectResult(request, config.SellerConfig)
	c.IndentedJSON(http.StatusOK, result)
}

func GetAllInspections(c *gin.Context) {
	result := models.QueryAllInspectResult(config.SellerConfig)
	c.IndentedJSON(http.StatusOK, result)
}
