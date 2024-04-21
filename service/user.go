package service

import (
	"context"
	"database/sql"

	"go.uber.org/zap"

	db "git.virjar.com/Junhiee/anilismei/database"
	"git.virjar.com/Junhiee/anilismei/database/models"
	"git.virjar.com/Junhiee/anilismei/pkg/log"
)

type UserService struct{}

var UserServer = new(UserService)

type User struct {
	UserID    int64
	UserName  string
	Email     string
	UserPwd   string
	AvatarUrl string
}

func (s *UserService) AddUser(u User) error {
	arg := models.AddUserParams{
		// UserID:    u.UserID,
		UserName:  u.UserName,
		Email:     u.Email,
		UserPwd:   u.UserPwd,
		AvatarUrl: sql.NullString{String: u.AvatarUrl, Valid: true},
	}
	err := db.G_QRY.AddUser(context.Background(), arg)

	if err != nil {
		log.ZLOG.Error("DB AddUser Err", zap.Error(err))
	}

	return err
}

func (s *UserService) UpdateUser(u User) error {
	arg := models.UpdateUserParams{
		UserID:    u.UserID,
		UserName:  u.UserName,
		Email:     u.Email,
		UserPwd:   u.UserPwd,
		AvatarUrl: sql.NullString{String: u.AvatarUrl, Valid: true},
	}

	err := db.G_QRY.UpdateUser(context.Background(), arg)
	if err != nil {
		log.ZLOG.Error("DB UpdateUser Err", zap.Error(err))
	}

	return err
}

func (s *UserService) DeleteUser(user_id int64) error {
	err := db.G_QRY.DeleteUser(context.Background(), user_id)
	if err != nil {
		log.ZLOG.Error("DB DeleteUser Err", zap.Error(err))
	}
	return err
}

func (s *UserService) GetUser(user_id int64) (models.User, error) {
	res, err := db.G_QRY.GetUser(context.Background(), user_id)
	if err != nil {
		log.ZLOG.Error("DB GetUser Err", zap.Error(err))
	}

	return res, err
}
