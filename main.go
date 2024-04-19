package main

import (
	"net/http"
	"time"

	db "git.virjar.com/Junhiee/anilismei/database"
	"git.virjar.com/Junhiee/anilismei/pkg/config"
	"git.virjar.com/Junhiee/anilismei/pkg/log"
	"git.virjar.com/Junhiee/anilismei/router"
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
