package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Opening"})
		})
		v1.POST("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Opening"})
		})
		v1.DELETE("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Opening"})
		})
		v1.PUT("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Opening"})
		})

	}

}
