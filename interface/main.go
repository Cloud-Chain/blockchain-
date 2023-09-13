package main

import (
	"github.com/gin-gonic/gin"
	"interface/config"
	"interface/router"
)

func main() {
	// Gin 라우터 초기화
	r := gin.Default()

	config.SellerConfig.Connect()
	config.BuyerConfig.Connect()
	config.InspectorConfig.Connect()

	// router.go에서 정의한 라우팅 설정을 가져옴
	router.SetupRouter(r)

	// 웹 서버 시작
	r.Run(":8080") // 8080 포트에서 웹 서버 실행
}
