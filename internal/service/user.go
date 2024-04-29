package service

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	models "github.com/Junhiee/anilismei/internal/models/sqlc"
	"github.com/Junhiee/anilismei/pkg/log"
)

type userService struct {
	db *models.Store
}

func NewUserService(db *models.Store) UserService {
	return &userService{db: db}
}

type User struct {
	UserID    int64
	UserName  string
	Email     string
	UserPwd   string
	AvatarUrl string
}

func (s *userService) AddUser(u User) error {
	arg := models.AddUserParams{
		// UserID:    u.UserID,
		UserName:  u.UserName,
		Email:     u.Email,
		UserPwd:   u.UserPwd,
		AvatarUrl: sql.NullString{String: u.AvatarUrl, Valid: true},
	}
	err := s.db.AddUser(context.Background(), arg)

	if err != nil {
		log.ZLOG.Error("DB AddUser Err", zap.Error(err))
	}

	return err
}

func (s *userService) UpdateUser(u User) error {
	arg := models.UpdateUserParams{
		UserID:    u.UserID,
		UserName:  u.UserName,
		Email:     u.Email,
		UserPwd:   u.UserPwd,
		AvatarUrl: sql.NullString{String: u.AvatarUrl, Valid: true},
	}

	err := s.db.UpdateUser(context.Background(), arg)
	if err != nil {
		log.ZLOG.Error("DB UpdateUser Err", zap.Error(err))
	}

	return err
}

func (s *userService) DeleteUser(user_id int64) error {
	err := s.db.DeleteUser(context.Background(), user_id)
	if err != nil {
		log.ZLOG.Error("DB DeleteUser Err", zap.Error(err))
	}
	return err
}

func (s *userService) GetUser(user_id int64) (models.User, error) {
	res, err := s.db.GetUser(context.Background(), user_id)
	if err != nil {
		log.ZLOG.Error("DB GetUser Err", zap.Error(err))
	}
	fmt.Println(res)
	return res, err
}
