package models

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dsn      = "root:123456@tcp(127.0.0.1:3306)/anilismei?parseTime=true"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Panic(err)
	}

	testDB.SetMaxOpenConns(64)
	testDB.SetMaxIdleConns(64)
	testDB.SetConnMaxLifetime(5 * time.Minute)
	testQueries = New(testDB)
	os.Exit(m.Run())
}
