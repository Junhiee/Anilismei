// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package models

import (
	"context"
)

type Querier interface {
	AddAnime(ctx context.Context, arg AddAnimeParams) error
	AddUser(ctx context.Context, arg AddUserParams) error
	DeleteUser(ctx context.Context, userID int64) error
	GetAnime(ctx context.Context, animeID int64) (Animation, error)
	GetListAnimes(ctx context.Context, arg GetListAnimesParams) ([]Animation, error)
	GetUser(ctx context.Context, userID int64) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
