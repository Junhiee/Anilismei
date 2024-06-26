// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: anime.sql

package models

import (
	"context"
	"database/sql"
	"strings"
)

const getAnimeByID = `-- name: GetAnimeByID :one
SELECT anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time, release_date, anime_status, rating FROM animations
WHERE anime_id = ?
`

func (q *Queries) GetAnimeByID(ctx context.Context, animeID int64) (Animation, error) {
	row := q.db.QueryRowContext(ctx, getAnimeByID, animeID)
	var i Animation
	err := row.Scan(
		&i.AnimeID,
		&i.GenreID,
		&i.StudioID,
		&i.Title,
		&i.Country,
		&i.ImageUrl,
		&i.Evaluate,
		&i.UpdateTime,
		&i.ReleaseDate,
		&i.AnimeStatus,
		&i.Rating,
	)
	return i, err
}

const getAnimeGenreID = `-- name: GetAnimeGenreID :one
SELECT genre_id FROM genres
WHERE genre_name = ?
`

func (q *Queries) GetAnimeGenreID(ctx context.Context, genreName string) (int32, error) {
	row := q.db.QueryRowContext(ctx, getAnimeGenreID, genreName)
	var genre_id int32
	err := row.Scan(&genre_id)
	return genre_id, err
}

const getAnimesByCountry = `-- name: GetAnimesByCountry :many
SELECT anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time, release_date, anime_status, rating FROM animations
WHERE country = ?
ORDER BY update_time DESC
LIMIT ? OFFSET ?
`

type GetAnimesByCountryParams struct {
	Country string `json:"country"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *Queries) GetAnimesByCountry(ctx context.Context, arg GetAnimesByCountryParams) ([]Animation, error) {
	rows, err := q.db.QueryContext(ctx, getAnimesByCountry, arg.Country, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Animation{}
	for rows.Next() {
		var i Animation
		if err := rows.Scan(
			&i.AnimeID,
			&i.GenreID,
			&i.StudioID,
			&i.Title,
			&i.Country,
			&i.ImageUrl,
			&i.Evaluate,
			&i.UpdateTime,
			&i.ReleaseDate,
			&i.AnimeStatus,
			&i.Rating,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAnimesByIDs = `-- name: GetAnimesByIDs :many
SELECT anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time, release_date, anime_status, rating FROM animations
WHERE anime_id IN (/*SLICE:ids*/?)
ORDER BY update_time DESC
`

func (q *Queries) GetAnimesByIDs(ctx context.Context, ids []int64) ([]Animation, error) {
	query := getAnimesByIDs
	var queryParams []interface{}
	if len(ids) > 0 {
		for _, v := range ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Animation{}
	for rows.Next() {
		var i Animation
		if err := rows.Scan(
			&i.AnimeID,
			&i.GenreID,
			&i.StudioID,
			&i.Title,
			&i.Country,
			&i.ImageUrl,
			&i.Evaluate,
			&i.UpdateTime,
			&i.ReleaseDate,
			&i.AnimeStatus,
			&i.Rating,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAnimesByRelease = `-- name: GetAnimesByRelease :many
SELECT anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time, release_date, anime_status, rating FROM animations
WHERE release_date = ?
ORDER BY update_time DESC
LIMIT ? OFFSET ?
`

type GetAnimesByReleaseParams struct {
	ReleaseDate sql.NullTime `json:"release_date"`
	Limit       int32        `json:"limit"`
	Offset      int32        `json:"offset"`
}

func (q *Queries) GetAnimesByRelease(ctx context.Context, arg GetAnimesByReleaseParams) ([]Animation, error) {
	rows, err := q.db.QueryContext(ctx, getAnimesByRelease, arg.ReleaseDate, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Animation{}
	for rows.Next() {
		var i Animation
		if err := rows.Scan(
			&i.AnimeID,
			&i.GenreID,
			&i.StudioID,
			&i.Title,
			&i.Country,
			&i.ImageUrl,
			&i.Evaluate,
			&i.UpdateTime,
			&i.ReleaseDate,
			&i.AnimeStatus,
			&i.Rating,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAnimesByType = `-- name: GetAnimesByType :many
SELECT anime_id FROM animation_genres
WHERE genre_id = ?
LIMIT ? OFFSET ?
`

type GetAnimesByTypeParams struct {
	GenreID sql.NullInt32 `json:"genre_id"`
	Limit   int32         `json:"limit"`
	Offset  int32         `json:"offset"`
}

func (q *Queries) GetAnimesByType(ctx context.Context, arg GetAnimesByTypeParams) ([]sql.NullInt64, error) {
	rows, err := q.db.QueryContext(ctx, getAnimesByType, arg.GenreID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []sql.NullInt64{}
	for rows.Next() {
		var anime_id sql.NullInt64
		if err := rows.Scan(&anime_id); err != nil {
			return nil, err
		}
		items = append(items, anime_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListAnimes = `-- name: GetListAnimes :many
SELECT anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time, release_date, anime_status, rating FROM animations
ORDER BY update_time DESC
LIMIT ? OFFSET ?
`

type GetListAnimesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListAnimes(ctx context.Context, arg GetListAnimesParams) ([]Animation, error) {
	rows, err := q.db.QueryContext(ctx, getListAnimes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Animation{}
	for rows.Next() {
		var i Animation
		if err := rows.Scan(
			&i.AnimeID,
			&i.GenreID,
			&i.StudioID,
			&i.Title,
			&i.Country,
			&i.ImageUrl,
			&i.Evaluate,
			&i.UpdateTime,
			&i.ReleaseDate,
			&i.AnimeStatus,
			&i.Rating,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAnime = `-- name: InsertAnime :exec
INSERT INTO animations(anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time,
release_date, anime_status, rating) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type InsertAnimeParams struct {
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

func (q *Queries) InsertAnime(ctx context.Context, arg InsertAnimeParams) error {
	_, err := q.db.ExecContext(ctx, insertAnime,
		arg.AnimeID,
		arg.GenreID,
		arg.StudioID,
		arg.Title,
		arg.Country,
		arg.ImageUrl,
		arg.Evaluate,
		arg.UpdateTime,
		arg.ReleaseDate,
		arg.AnimeStatus,
		arg.Rating,
	)
	return err
}
