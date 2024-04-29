package service

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/zap"

	models "github.com/Junhiee/anilismei/internal/models/sqlc"
	"github.com/Junhiee/anilismei/pkg/log"
)



type animeService struct {
	db *models.Store
}

func NewAnimeService(db *models.Store) AnimeService {
	return &animeService{db: db}
}

type Animation struct {
	AnimeID     int64
	GenreID     int32
	StudioID    int32
	Title       string
	Country     string
	ImageUrl    string
	Evaluate    string
	UpdateTime  time.Time
	ReleaseDate time.Time
	AnimeStatus string
	Rating      float64
}

// 暂时用不上，业务逻辑以后修改
func (s *animeService) AddAnime(a Animation) error {
	anime := models.InsertAnimeParams{
		AnimeID:     a.AnimeID,
		GenreID:     a.GenreID,
		StudioID:    a.StudioID,
		Title:       a.Title,
		ImageUrl:    sql.NullString{String: a.ImageUrl, Valid: true},
		Evaluate:    sql.NullString{String: a.Evaluate, Valid: true},
		UpdateTime:  sql.NullTime{Time: a.UpdateTime, Valid: true},
		ReleaseDate: sql.NullTime{Time: a.ReleaseDate, Valid: true},
		AnimeStatus: models.NullAnimationsAnimeStatus{AnimationsAnimeStatus: models.AnimationsAnimeStatus(a.AnimeStatus), Valid: true},
		Rating:      sql.NullFloat64{Float64: a.Rating, Valid: true},
	}

	err := s.db.InsertAnime(context.Background(), anime)
	if err != nil {
		log.ZLOG.Error("DB insert anime err", zap.Error(err))
	}

	return err
}

// 通过 anime_id 获取结果
func (s *animeService) GetAnimeByID(anime_id int64) (models.Animation, error) {
	res, err := s.db.GetAnimeByID(context.Background(), anime_id)
	if err != nil {
		log.ZLOG.Error("GetAnimeByID err", zap.Error(err))
	}
	return res, err
}

// 筛选出按动画类型分类的结果
func (s *animeService) GetAnimesByCountry(country string, limit int32, offset int32) ([]models.Animation, error) {

	params := models.GetAnimesByCountryParams{
		Country: country,
		Limit:   limit,
		Offset:  offset,
	}

	res, err := s.db.GetAnimesByCountry(context.Background(), params)

	if err != nil {
		log.ZLOG.Error("GetAnimesByCountry err", zap.Error(err))
	}

	return res, err

}

// 获取动画列表默认按更新日期排序
func (s *animeService) GetListAnimes(limit int32, offset int32) ([]models.Animation, error) {
	params := models.GetListAnimesParams{
		Limit:  limit,
		Offset: offset,
	}
	res, err := s.db.GetListAnimes(context.Background(), params)

	if err != nil {
		log.ZLOG.Error("GetListAnimes err", zap.Error(err))
	}
	return res, err
}

// 根据推出日期获得数据
func (s *animeService) GetAnimesByRelease() ([]models.Animation, error) {
	params := models.GetAnimesByReleaseParams{}
	res, err := s.db.GetAnimesByRelease(context.Background(), params)

	if err != nil {
		log.ZLOG.Error("GetAnimesByRelease err", zap.Error(err))
	}
	return res, err
}

// 根据动画类型获得数据
func (s *animeService) GetAnimesByType() ([]models.AnimationGenre, error) {
	params := models.GetAnimesByTypeParams{}
	res, err := s.db.GetAnimesByType(context.Background(), params)

	if err != nil {
		log.ZLOG.Error("GetAnimesByType err", zap.Error(err))
	}
	return res, err
}
