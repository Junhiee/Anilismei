package router

import (
	"time"

	"git.virjar.com/Junhiee/anilismei/pkg/log"
	"git.virjar.com/Junhiee/anilismei/router/api/beta"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

type RoterGroup struct {
	beta.AnimeRouter
}

var RoterGroups = new(RoterGroup)

func Routers() *gin.Engine {
	router := gin.New()

	router.Use(ginzap.Ginzap(log.ZLOG, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(log.ZLOG, true))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api_beta := router.Group("/api/beta/anime")
	{
		api_beta.GET(":anime_id", RoterGroups.GetAnime)
		api_beta.GET("/", RoterGroups.GetListAnimes)
		api_beta.POST("/", RoterGroups.AddAnime)
		api_beta.PUT("/:id", RoterGroups.UpdateAnime)
		api_beta.DELETE("/:id", RoterGroups.DeleteAnime)
	}

	return router
}
