package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	_db "github.com/Junhiee/anilismei/internal/models/sqlc"
	"github.com/Junhiee/anilismei/pkg/log"
)

// var G_QRY *_db.Store

func InitMysql() *_db.Store {
	dsn := "root:123456@tcp(127.0.0.1:3306)/anilismei?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.ZLOG.Panic("mysql setup err: ", zap.Error(err))
		return nil
	}

	db.SetMaxOpenConns(64)
	db.SetMaxIdleConns(64)
	db.SetConnMaxLifetime(5 * time.Minute)
	stroe := _db.NewStore(db)
	return stroe
}