package service

import (
	"time"

	models "github.com/Junhiee/anilismei/internal/models/sqlc"
)

type AnimeService interface {
	AddAnime(a Animation) error
	GetAnimeByID(anime_id int64) (models.Animation, error)
	GetAnimesByCountry(country string, limit, offset int32) ([]models.Animation, error)
	GetListAnimes(limit int32, offset int32) ([]models.Animation, error)
	GetAnimesByRelease(release_date time.Time, limit, offset int32) ([]models.Animation, error)
	GetAnimesByType(genre_name string, limit, offset int32) ([]models.Animation, error)
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
