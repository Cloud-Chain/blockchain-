package handlers

import (
	"interface/config"
	"interface/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
	result := models.InspectResult(request, config.InspectorConfig)
	c.IndentedJSON(http.StatusOK, result)
}

func FindInspection(c *gin.Context) {
	// request, err := strconv.ParseInt(c.Query("id"), 10, 64)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
	// }
	request := c.Query("id")
	result := models.QueryInspectResult(request, config.InspectorConfig)
	c.IndentedJSON(http.StatusOK, result)
}

func GetAllInspections(c *gin.Context) {
	result := models.QueryAllInspections(config.InspectorConfig)
	c.IndentedJSON(http.StatusOK, result)
}
