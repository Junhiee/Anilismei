package initialize

import (
	"database/sql"
	"fmt"
	"time"

	datebase "git.virjar.com/Junhiee/anilismei/database"
	"git.virjar.com/Junhiee/anilismei/global"
	_ "github.com/go-sql-driver/mysql"
)

func SetupMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/anilisme"
	if db, err := sql.Open("mysql", dsn); err != nil {
		fmt.Println(err)
	} else {
		db.SetMaxOpenConns(64)
		db.SetMaxIdleConns(64)
		db.SetConnMaxLifetime(5 * time.Minute)
		global.G_QRY = datebase.NewStore(db)
	}
}
