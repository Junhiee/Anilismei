package models

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Junhiee/anilismei/tools"
)

func GenAnimeParams(id int, n int, isEmpty bool) []InsertAnimeParams {
	var AnimeParams []InsertAnimeParams
	for i := id; i < (id + n); i++ {

		params := InsertAnimeParams{
			AnimeID:     int64(i),
			GenreID:     int32(i),
			StudioID:    int32(i),
			Title:       tools.RandomString(10),
			Country:     tools.RandomString(5),
			ImageUrl:    sql.NullString{String: tools.RandomString(5), Valid: true},
			Evaluate:    sql.NullString{String: tools.RandomString(25), Valid: true},
			UpdateTime:  sql.NullTime{Time: time.Now(), Valid: true},
			ReleaseDate: sql.NullTime{Time: time.Now(), Valid: true},
			AnimeStatus: NullAnimationsAnimeStatus{AnimationsAnimeStatus: "coming soon", Valid: true},
			Rating:      sql.NullFloat64{Float64: 9.2, Valid: true},
		}

		if !isEmpty {
			params.AnimeID = 0
		}

		AnimeParams = append(AnimeParams, params)
	}

	return AnimeParams
}

func TestAddAnime(t *testing.T) {
	var animes []InsertAnimeParams

	rs, err := testQueries.GetAnimeByID(context.Background(), 90000001)
	require.NoError(t, err)
	if !assert.NotEmpty(t, rs) {
		animes = GenAnimeParams(90000001, 5, true)
	} else {
		animes = GenAnimeParams(90000001, 5, false)
	}

	for _, anime := range animes {
		err := testQueries.InsertAnime(context.Background(), anime)
		require.NoError(t, err)
	}

}

func TestGetAnimeByID(t *testing.T) {
	res, err := testQueries.GetAnimeByID(context.Background(), 90000001)

	require.NotEmpty(t, res)
	require.NoError(t, err)
}

// TODO 4. GetListAnimes test
