package models

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddAnime(t *testing.T) {
	anime := AddAnimeParams{
		AnimeID:     100004,
		GenreID:     100004,
		StudioID:    100004,
		Title:       "怪兽8号",
		Country:     "日本",
		ImageUrl:    sql.NullString{String: "https://127.0.0.1:8081", Valid: true},
		Evaluate:    sql.NullString{String: "在日本这一“怪兽大国”，人们的日常生活遭受着怪兽的无情破坏", Valid: true},
		UpdateTime: sql.NullTime{Time: time.Now(), Valid: true},
		ReleaseDate: sql.NullTime{Time: time.Now(), Valid: true},
		AnimeStatus: NullAnimationsAnimeStatus{AnimationsAnimeStatus: AnimationsAnimeStatus("completed"), Valid: true},
		Rating:      sql.NullFloat64{Float64: 9.0, Valid: true},
	}

	err := testQueries.AddAnime(context.Background(), anime)

	require.NoError(t, err)

}

func TestGetAnime(t *testing.T) {
	res, err := testQueries.GetAnime(context.Background(), 100001)

	require.NotEmpty(t, res)
	require.NoError(t, err)
}

type GetListAnimesParamsServer struct {
	*GetListAnimesParams
	Country string
	Genre   string
	Date    time.Time
	Sort    string
}

func TestGetListAnimes(t *testing.T) {

	params := GetListAnimesParams{
		Limit:  5,
		Offset: 0,
	}

	server_params := GetListAnimesParamsServer{
		Country: "日本",
		Genre:   "action",
		Date:    time.Now(),
		Sort:    "hot",
	}

	animes, err := testQueries.GetListAnimes(context.Background(), params)

	require.NoError(t, err)
	require.Equal(t, len(animes), 5)

	for _, anime := range animes {
		if server_params.Country != "" && server_params.Country == anime.Country {
			log.Println(anime)
		}

	}

}
