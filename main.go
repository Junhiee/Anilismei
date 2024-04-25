package main

import (
	"net/http"
	"time"

	"github.com/Junhiee/anilismei/internal/router"
	"github.com/Junhiee/anilismei/pkg/config"
	db "github.com/Junhiee/anilismei/pkg/db"
	"github.com/Junhiee/anilismei/pkg/log"
)

func main() {

	db.InitMysql()     // 初始化数据库
	log.InitLogger()    // 初始化日志服务
	config.InitConfig() // 初始化配置文件

	defer log.ZLOG.Sync()
	r := router.InitRouters()

	s := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
