package initialize

import (
	"database/sql"
	"time"

	// "git.virjar.com/Junhiee/anilismei/database/models"
	_ "github.com/go-sql-driver/mysql"
)

// TODO init mysql
func MysqlC() *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/onlinedb"
	if db, err := sql.Open("mysql", dsn); err != nil {
		return nil
	} else {
		db.SetMaxOpenConns(64)
		db.SetMaxIdleConns(64)
		db.SetConnMaxLifetime(5 * time.Minute)

		// q := models.New(db)
		return db
	}
}
