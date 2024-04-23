package main

import (
	"net/http"
	"time"

	"git.virjar.com/Junhiee/anilismei/internal/router"
	"git.virjar.com/Junhiee/anilismei/pkg/config"
	db "git.virjar.com/Junhiee/anilismei/pkg/db"
	"git.virjar.com/Junhiee/anilismei/pkg/log"
)

func main() {

	db.SetupMysql()     // 初始化数据库
	log.InitLogger()    // 初始化日志服务
	config.InitConfig() // 初始化配置文件

	r := router.Routers()
	s := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
