package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	_db "git.virjar.com/Junhiee/anilismei/internal/models/sqlc"
	"git.virjar.com/Junhiee/anilismei/pkg/log"
)

var G_QRY *_db.Store

func SetupMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/anilismei?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.ZLOG.Error("mysql setup err: ", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(64)
	db.SetMaxIdleConns(64)
	db.SetConnMaxLifetime(5 * time.Minute)
	G_QRY = _db.NewStore(db)
}
