package handlers

import (
	"github.com/gin-gonic/gin"
	"interface/config"
	"interface/models"
	"net/http"
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
	result := models.InspectRequest(request.ID, request.BasicInfo, config.SellerConfig)
	c.IndentedJSON(http.StatusCreated, result)
}

func ExecuteInspection(c *gin.Context) {
	var request models.Inspection
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.InspectRequest(request.ID, request.BasicInfo, config.SellerConfig)
	c.IndentedJSON(http.StatusCreated, result)
}
