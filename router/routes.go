package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "GET Vaga",
			})
		})

		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "POST Vagas",
			})
		})

		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "DELETE Vagas",
			})
		})

		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "PUT Vagas",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "GET Vagas",
			})
		})
	}
}
