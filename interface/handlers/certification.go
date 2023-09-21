package handlers

import (
	"github.com/gin-gonic/gin"
	"interface/models"
	"net/http"
)

// CA 인증서 발급
func Enroll(c *gin.Context) {
	var request models.CertRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.EnrollUser(request)
	c.IndentedJSON(http.StatusCreated, result)
}
