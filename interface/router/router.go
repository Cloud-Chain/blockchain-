package router

import (
	"github.com/gin-gonic/gin"
	"interface/config"
	"interface/handlers"
	"interface/models"
)

func SetupRouter(r *gin.Engine) {

	models.TransactionInitLedger(config.SellerConfig)
	models.InspectionInitLedger(config.SellerConfig)

	//거래 라우팅 설정
	tx := r.Group("/tx")
	{
		tx.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
	}

	//검수 라우팅 설정
	ix := r.Group("/ix")
	{
		ix.GET("/inspect", handlers.FindInspection)
		ix.POST("/inspect", handlers.RequestInspection)
		ix.PATCH("/inspect", handlers.ExecuteInspection)
		ix.GET("/", handlers.GetAllInspections)
	}

	// 기타 라우팅 설정
	// 루트 URL에 대한 핸들러 등록
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "transaction",
		})
	})

	// 다른 URL 경로에 대한 핸들러 등록
	r.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "About Us",
		})
	})

}
