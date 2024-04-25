// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: anime.sql

package models

import (
	"context"
	"database/sql"
)

const addAnime = `-- name: AddAnime :exec
INSERT INTO animations(anime_id, genre_id, studio_id, title, evaluate , release_date, anime_status, rating)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type AddAnimeParams struct {
	AnimeID     int64                     `json:"anime_id"`
	GenreID     int32                     `json:"genre_id"`
	StudioID    int32                     `json:"studio_id"`
	Title       string                    `json:"title"`
	Evaluate    string                    `json:"evaluate"`
	ReleaseDate sql.NullTime              `json:"release_date"`
	AnimeStatus NullAnimationsAnimeStatus `json:"anime_status"`
	Rating      sql.NullFloat64           `json:"rating"`
}

func (q *Queries) AddAnime(ctx context.Context, arg AddAnimeParams) error {
	_, err := q.db.ExecContext(ctx, addAnime,
		arg.AnimeID,
		arg.GenreID,
		arg.StudioID,
		arg.Title,
		arg.Evaluate,
		arg.ReleaseDate,
		arg.AnimeStatus,
		arg.Rating,
	)
	return err
}

const deleteAnime = `-- name: DeleteAnime :exec
DELETE FROM animations
WHERE anime_id = ?
`

func (q *Queries) DeleteAnime(ctx context.Context, animeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAnime, animeID)
	return err
}

const getAnime = `-- name: GetAnime :one
SELECT anime_id, genre_id, studio_id, title, evaluate, release_date, anime_status, rating FROM animations
WHERE anime_id = ?
`

func (q *Queries) GetAnime(ctx context.Context, animeID int64) (Animation, error) {
	row := q.db.QueryRowContext(ctx, getAnime, animeID)
	var i Animation
	err := row.Scan(
		&i.AnimeID,
		&i.GenreID,
		&i.StudioID,
		&i.Title,
		&i.Evaluate,
		&i.ReleaseDate,
		&i.AnimeStatus,
		&i.Rating,
	)
	return i, err
}

const getListAnimes = `-- name: GetListAnimes :many
SELECT anime_id, genre_id, studio_id, title, evaluate, release_date, anime_status, rating FROM animations 
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
			&i.Evaluate,
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

const updateAnime = `-- name: UpdateAnime :exec
UPDATE animations
SET genre_id = ?, studio_id = ?, title = ?, evaluate = ?, release_date = ?, anime_status = ?, rating = ?
WHERE anime_id = ?
`

type UpdateAnimeParams struct {
	GenreID     int32                     `json:"genre_id"`
	StudioID    int32                     `json:"studio_id"`
	Title       string                    `json:"title"`
	Evaluate    string                    `json:"evaluate"`
	ReleaseDate sql.NullTime              `json:"release_date"`
	AnimeStatus NullAnimationsAnimeStatus `json:"anime_status"`
	Rating      sql.NullFloat64           `json:"rating"`
	AnimeID     int64                     `json:"anime_id"`
}

func (q *Queries) UpdateAnime(ctx context.Context, arg UpdateAnimeParams) error {
	_, err := q.db.ExecContext(ctx, updateAnime,
		arg.GenreID,
		arg.StudioID,
		arg.Title,
		arg.Evaluate,
		arg.ReleaseDate,
		arg.AnimeStatus,
		arg.Rating,
		arg.AnimeID,
	)
	return err
}
