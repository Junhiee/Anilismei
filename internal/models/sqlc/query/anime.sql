-- name: GetAnimeByID :one
SELECT * FROM animations
WHERE anime_id = ?;

-- name: GetAnimesByIDs :many
SELECT * FROM animations
WHERE anime_id IN (sqlc.slice('ids'))
ORDER BY update_time DESC;

-- name: GetListAnimes :many
SELECT * FROM animations
ORDER BY update_time DESC
LIMIT ? OFFSET ?;

-- name: GetAnimesByCountry :many
SELECT * FROM animations
WHERE country = ?
ORDER BY update_time DESC
LIMIT ? OFFSET ?;

-- name: GetAnimesByRelease :many
SELECT * FROM animations
WHERE release_date = ?
ORDER BY update_time DESC;

-- name: GetAnimesByType :many
SELECT * FROM animation_genres
WHERE genre_id = ?;

-- name: InsertAnime :exec
INSERT INTO animations(anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time,
release_date, anime_status, rating) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);