-- name: GetAnime :one
SELECT * FROM animations
WHERE anime_id = ?;

-- name: GetListAnimes :many
SELECT * FROM animations
ORDER BY update_time
LIMIT ? OFFSET ?;

-- name: AddAnime :exec
INSERT INTO animations(anime_id, genre_id, studio_id, title, country, image_url, evaluate, update_time, release_date, anime_status, rating)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);