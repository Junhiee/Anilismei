package service

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/zap"

	models "github.com/Junhiee/anilismei/internal/models/sqlc"
	db "github.com/Junhiee/anilismei/pkg/db"
	"github.com/Junhiee/anilismei/pkg/log"
)

type AnimeService struct{}

var AnimeServer = new(AnimeService)

type Animation struct {
	AnimeID     int64
	Title       string
	Evaluate    string
	GenreID     int32
	ReleaseDate time.Time
	StudioID    int32
	AnimeStatus string
	Rating      float64
}

func (s *AnimeService) AddAnime(a Animation) error {
	anime := models.AddAnimeParams{
		AnimeID:     a.AnimeID,
		Title:       a.Title,
		Evaluate:    a.Evaluate,
		GenreID:     a.GenreID,
		ReleaseDate: sql.NullTime{Time: a.ReleaseDate, Valid: true},
		StudioID:    a.StudioID,
		AnimeStatus: models.NullAnimationsAnimeStatus{AnimationsAnimeStatus: models.AnimationsAnimeStatus(a.AnimeStatus), Valid: true},
		Rating:      sql.NullFloat64{Float64: a.Rating, Valid: true},
	}

	err := db.G_QRY.AddAnime(context.Background(), anime)
	if err != nil {
		log.ZLOG.Error("DB insert anime err", zap.Error(err))
	}

	return err
}

func (s *AnimeService) GetAnime(anime_id int64) (*models.Animation, error) {

	res, err := db.G_QRY.GetAnime(context.Background(), anime_id)
	if err != nil {
		log.ZLOG.Error("Get anime err", zap.Error(err))
		return nil, err
	}
	return &res, err
}

func (s *AnimeService) GetListAnimes(limit int32, offset int32) ([]models.Animation, error) {
	params := models.GetListAnimesParams{
		Limit:  limit,
		Offset: offset,
	}
	res, err := db.G_QRY.GetListAnimes(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *AnimeService) UpdateAnime(a Animation) error {
	params := models.UpdateAnimeParams{
		AnimeID:     a.AnimeID,
		Title:       a.Title,
		Evaluate:    a.Evaluate,
		GenreID:     a.GenreID,
		ReleaseDate: sql.NullTime{Time: a.ReleaseDate, Valid: true},
		StudioID:    a.StudioID,
		AnimeStatus: models.NullAnimationsAnimeStatus{AnimationsAnimeStatus: models.AnimationsAnimeStatus(a.AnimeStatus), Valid: true},
		Rating:      sql.NullFloat64{Float64: a.Rating, Valid: true},
	}

	err := db.G_QRY.UpdateAnime(context.Background(), params)

	return err
}

func (s *AnimeService) DeleteAnime(aid int64) error {
	err := db.G_QRY.DeleteAnime(context.Background(), aid)
	return err
}
