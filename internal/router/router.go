package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"github.com/Junhiee/anilismei/internal/api/beta"
	"github.com/Junhiee/anilismei/internal/service"
	"github.com/Junhiee/anilismei/pkg/log"
)

// TODO 根据人气排序 -- 用缓存来做

// TODO 筛选出按动画类型分类的结果

// TODO 筛选出按推出日期分类的结果

type RouterGroup struct {
	AnimeR *beta.AnimeRouter
	UserR  *beta.UserRouter
}

func NewRouterGroup(service *service.Service) *RouterGroup {
	return &RouterGroup{
		AnimeR: beta.NewAnimeRouter(service.AnimeSrv),
		UserR:  beta.NewUserRouter(service.UserSrv),
	}
}

func InitRouters(service *service.Service) *gin.Engine {
	beta := NewRouterGroup(service)
	router := gin.New()
	router.Use(ginzap.Ginzap(log.ZLOG, time.RFC3339, true))
	// router.Use(ginzap.RecoveryWithZap(log.ZLOG, true))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	anime_beta := router.Group("/api/beta/anime")
	{
		anime_beta.GET("", beta.AnimeR.GetListAnimes)
		anime_beta.GET(":anime_id", beta.AnimeR.GetAnime)

	}

	user_beta := router.Group("/api/beta/user")
	{
		user_beta.GET("/:user_id", beta.UserR.GetUser)
		user_beta.POST("/", beta.UserR.AddUser)

	}
	return router
}
