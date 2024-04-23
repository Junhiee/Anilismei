package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"git.virjar.com/Junhiee/anilismei/internal/router/api/beta"
	"git.virjar.com/Junhiee/anilismei/pkg/log"
)

type RoterGroup struct {
	beta.AnimeRouter
	beta.UserRouter
}

var RoterGroups = new(RoterGroup)

func Routers() *gin.Engine {
	router := gin.New()

	router.Use(ginzap.Ginzap(log.ZLOG, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(log.ZLOG, true))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	anime_beta := router.Group("/api/beta/anime")
	{
		anime_beta.GET(":anime_id", RoterGroups.GetAnime)
		anime_beta.GET("/", RoterGroups.GetListAnimes)
		anime_beta.POST("/", RoterGroups.AddAnime)
		anime_beta.PUT("/:id", RoterGroups.UpdateAnime)
		anime_beta.DELETE("/:id", RoterGroups.DeleteAnime)
	}

	user_beta := router.Group("/api/beta/user")
	{
		user_beta.GET("/:user_id", RoterGroups.GetUser)
		user_beta.POST("/", RoterGroups.AddUser)

	}
	return router
}
