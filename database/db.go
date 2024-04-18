package database

import (
	"database/sql"
	"time"

	"git.virjar.com/Junhiee/anilismei/database/models"
	"git.virjar.com/Junhiee/anilismei/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var G_QRY *models.Store

func SetupMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/anilisme?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.ZLOG.Error("mysql setup err: ", zap.Error(err))
		return
	}

	
	db.SetMaxOpenConns(64)
	db.SetMaxIdleConns(64)
	db.SetConnMaxLifetime(5 * time.Minute)
	G_QRY = models.NewStore(db)
}
