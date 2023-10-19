package router

import (
	"github.com/gin-gonic/gin"
	"interface/handlers"
)

func SetupRouter(r *gin.Engine) {

	// models.TransactionInitLedger(config.SellerConfig)
	// models.InspectionInitLedger(config.SellerConfig)

	//거래 라우팅 설정
	tx := r.Group("/tx")
	{
		tx.POST("/sell", handlers.SellVehicle)
		tx.POST("/buy", handlers.BuyVehicle)
		tx.PATCH("/seller/compromise", handlers.SellerCompromiseTransaction)
		tx.PATCH("/buyer/compromise", handlers.BuyerCompromiseTransaction)
		tx.GET("/id", handlers.ReadTransaction)
		tx.GET("/user", handlers.QueryTransactionsByUser)
		tx.GET("/vehicle", handlers.QueryTransactionsByVehicle)
		tx.GET("/", handlers.QueryAllTransactions)
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
	cert := r.Group("/cert")
	{

		cert.POST("/enroll", handlers.Enroll) // 인증서 발급
		cert.POST("/re-enroll")               // 인증서 재발급

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
