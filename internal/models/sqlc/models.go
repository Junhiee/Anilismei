// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type AnimationsAnimeStatus string

const (
	AnimationsAnimeStatusComingsoon AnimationsAnimeStatus = "coming soon"
	AnimationsAnimeStatusAiring     AnimationsAnimeStatus = "airing"
	AnimationsAnimeStatusCompleted  AnimationsAnimeStatus = "completed"
	AnimationsAnimeStatusPaused     AnimationsAnimeStatus = "paused"
)

func (e *AnimationsAnimeStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AnimationsAnimeStatus(s)
	case string:
		*e = AnimationsAnimeStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AnimationsAnimeStatus: %T", src)
	}
	return nil
}

type NullAnimationsAnimeStatus struct {
	AnimationsAnimeStatus AnimationsAnimeStatus `json:"animations_anime_status"`
	Valid                 bool                  `json:"valid"` // Valid is true if AnimationsAnimeStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAnimationsAnimeStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AnimationsAnimeStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AnimationsAnimeStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAnimationsAnimeStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AnimationsAnimeStatus), nil
}

type Animation struct {
	AnimeID     int64                     `json:"anime_id"`
	GenreID     int32                     `json:"genre_id"`
	StudioID    int32                     `json:"studio_id"`
	Title       string                    `json:"title"`
	Country     string                    `json:"country"`
	ImageUrl    sql.NullString            `json:"image_url"`
	Evaluate    sql.NullString            `json:"evaluate"`
	UpdateTime  sql.NullTime              `json:"update_time"`
	ReleaseDate sql.NullTime              `json:"release_date"`
	AnimeStatus NullAnimationsAnimeStatus `json:"anime_status"`
	Rating      sql.NullFloat64           `json:"rating"`
}

type AnimationGenre struct {
	AnimeID sql.NullInt32 `json:"anime_id"`
	GenreID sql.NullInt32 `json:"genre_id"`
}

type Comment struct {
	CommentID   int64     `json:"comment_id"`
	AnimeID     int64     `json:"anime_id"`
	CommentText string    `json:"comment_text"`
	CommentDate time.Time `json:"comment_date"`
}

type Genre struct {
	GenreID   int32  `json:"genre_id"`
	GenreName string `json:"genre_name"`
}

type Studio struct {
	StudioID    int32  `json:"studio_id"`
	StudioName  string `json:"studio_name"`
	StudioStaff string `json:"studio_staff"`
}

type Subscription struct {
	SubscriptionID   int32     `json:"subscription_id"`
	UserID           int64     `json:"user_id"`
	AnimeID          int64     `json:"anime_id"`
	SubscriptionDate time.Time `json:"subscription_date"`
}

type User struct {
	UserID    int64          `json:"user_id"`
	UserName  string         `json:"user_name"`
	Email     string         `json:"email"`
	UserPwd   string         `json:"user_pwd"`
	AvatarUrl sql.NullString `json:"avatar_url"`
}
