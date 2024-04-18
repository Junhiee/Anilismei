package service

import (
	"context"
	"database/sql"
	"time"

	"git.virjar.com/Junhiee/anilismei/database/models"
	"git.virjar.com/Junhiee/anilismei/global"
	"go.uber.org/zap"
)

type AnimeService struct{}

// var AnimeServer = new(AnimeService)

type Animation struct {
	AnimeID     int64
	Title       string
	Evaluate    string
	GenreID     int32
	ReleaseDate time.Time
	StudioID    int32
	AnimeStatus string
	Rating      string
}

func (s *AnimeService) Add(a Animation) error {
	anime := models.AddAnimeParams{
		AnimeID:     a.AnimeID,
		Title:       a.Title,
		Evaluate:    a.Evaluate,
		GenreID:     a.GenreID,
		ReleaseDate: sql.NullTime{Time: a.ReleaseDate, Valid: true},
		StudioID:    a.StudioID,
		AnimeStatus: models.NullAnimationsAnimeStatus{AnimationsAnimeStatus: models.AnimationsAnimeStatus(a.AnimeStatus), Valid: true},
		Rating:      sql.NullString{String: a.Rating, Valid: true},
	}

	err := global.G_QRY.AddAnime(context.Background(), anime)
	if err != nil {
		global.ZLOG.Error("Add anime err:", zap.Error(err))
	}

	return err
}

func (s *AnimeService) Get(anime_id int64) (*models.Animation, error) {

	res, err := global.G_QRY.GetAnime(context.Background(), anime_id)
	if err != nil {
		global.ZLOG.Error("Get anime err:", zap.Error(err))
		return nil, err
	}
	return &res, err
}

func (s *AnimeService) GetList(limit int32, offset int32) ([]models.Animation, error) {
	params := models.GetListAnimesParams{
		Limit:  limit,
		Offset: offset,
	}
	res, err := global.G_QRY.GetListAnimes(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *AnimeService) Update(a Animation) error {
	params := models.UpdateAnimeParams{
		AnimeID:     a.AnimeID,
		Title:       a.Title,
		Evaluate:    a.Evaluate,
		GenreID:     a.GenreID,
		ReleaseDate: sql.NullTime{Time: a.ReleaseDate, Valid: true},
		StudioID:    a.StudioID,
		AnimeStatus: models.NullAnimationsAnimeStatus{AnimationsAnimeStatus: models.AnimationsAnimeStatus(a.AnimeStatus), Valid: true},
		Rating:      sql.NullString{String: a.Rating, Valid: true},
	}

	err := global.G_QRY.UpdateAnime(context.Background(), params)
	
	return err
}

func (s *AnimeService) Delete(aid int64) error {
	err := global.G_QRY.DeleteAnime(context.Background(), aid)
	return err
}
