package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interface/models"
	"net/http"
	"strings"
)

// CA 인증서 발급
func Enroll(c *gin.Context) {
	var request models.CertRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	result := models.EnrollUser(request)
	// 개행 문자열을 정확하게 변환
	convertedResult := strings.ReplaceAll(string(result), "\n", "")
	fmt.Println(result)
	fmt.Println(convertedResult)
	c.Header("CA-User", result)
	c.IndentedJSON(http.StatusCreated, result)
}
