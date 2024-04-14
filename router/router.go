package router

import (
	"git.virjar.com/Junhiee/anilismei/router/api/beta"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	
	api_beta := router.Group("/api/beta")
	{
		api_beta.GET("/anime", beta.GetAnime)
		api_beta.POST("/anime", beta.AddAnime)
		api_beta.PUT("/anime:id", beta.UpdateAnime)
		api_beta.DELETE("/anime:id", beta.DeleteAnime)
	}

	return router
}
