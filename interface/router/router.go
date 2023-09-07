package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// 루트 URL에 대한 핸들러 등록
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 다른 URL 경로에 대한 핸들러 등록
	r.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "About Us",
		})
	})

	// 기타 라우팅 설정을 추가할 수 있습니다.
}
