package service

import (
	models "github.com/Junhiee/anilismei/internal/models/sqlc"
)

type AnimeService interface {
	AddAnime(a Animation) error
	GetAnimeByID(anime_id int64) (models.Animation, error)
	GetAnimesByCountry(country string, limit int32, offset int32) ([]models.Animation, error)
	GetListAnimes(limit int32, offset int32) ([]models.Animation, error)
	GetAnimesByRelease() ([]models.Animation, error)
	GetAnimesByType() ([]models.AnimationGenre, error)
}

type UserService interface {
	GetUser(user_id int64) (models.User, error)
	AddUser(u User) error
}

type Service struct {
	AnimeSrv AnimeService
	UserSrv  UserService
}

func NewService(db *models.Store) *Service {
	return &Service{
		AnimeSrv: NewAnimeService(db),
		UserSrv:  NewUserService(db),
	}
}
