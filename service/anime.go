package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"git.virjar.com/Junhiee/anilismei/database/models"
	"git.virjar.com/Junhiee/anilismei/global"
)

func Add() (err error) {
	anime := models.AddAnimeParams{
		AnimeID:     10004,
		Title:       "Test Anime",
		Evaluate:    "Great",
		GenreID:     10003,
		ReleaseDate: sql.NullTime{Time: time.Now(), Valid: true},
		StudioID:    10004,
		AnimeStatus: models.NullAnimationsAnimeStatus{AnimationsAnimeStatus: "completed", Valid: true},
		Rating:      sql.NullString{String: "6", Valid: true},
	}

	err = global.G_QRY.AddAnime(context.Background(), anime)
	if err != nil {
		fmt.Println(err)
	}

	return err

}
