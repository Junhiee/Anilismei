package datebase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"git.virjar.com/Junhiee/anilismei/database/models"
	_ "github.com/go-sql-driver/mysql"
)

var q *Store

func SetupMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/anilisme"
	if db, err := sql.Open("mysql", dsn); err != nil {
		fmt.Println(err)
	} else {
		db.SetMaxOpenConns(64)
		db.SetMaxIdleConns(64)
		db.SetConnMaxLifetime(5 * time.Minute)
		q = NewStore(db)
	}

}

func TestAdd(t *testing.T) {
	SetupMysql()
	anime := models.AddAnimeParams{
		AnimeID:     10001,
		Title:       "Test Anime",
		Evaluate:    "Great",
		GenreID:     10001,
		ReleaseDate: sql.NullTime{Time: time.Now(), Valid: true},
		StudioID:    10001,
		AnimeStatus: models.NullAnimationsAnimeStatus{AnimationsAnimeStatus: "completed", Valid: true},
		Rating:      sql.NullString{String: "9", Valid: true},
	}
	err := q.AddAnime(context.Background(), anime)
	if err != nil {
		fmt.Println(err)
	}

}
