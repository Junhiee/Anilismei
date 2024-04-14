-- name: GetListAnimes :many
SELECT * FROM animations
ORDER BY rating;

-- name: GetAnime :one
SELECT * FROM animations
WHERE anime_id = ?;

-- name: AddAnime :exec
INSERT INTO animations(anime_id, title, evaluate, genre_id, release_date, studio_id, anime_status, rating)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateAnime :exec
UPDATE animations
SET title = ?, evaluate = ?, genre_id = ?, release_date = ?, studio_id = ?, anime_status = ?, rating = ?
WHERE anime_id = ?;

-- name: DeleteAnime :exec
DELETE FROM animations
WHERE anime_id = ?;