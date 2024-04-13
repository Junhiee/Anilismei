package router

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return router
}
