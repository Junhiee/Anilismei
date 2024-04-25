-- name: GetListAnimes :many
SELECT * FROM animations 
LIMIT ? OFFSET ?;

-- name: GetAnime :one
SELECT * FROM animations
WHERE anime_id = ?;

-- name: AddAnime :exec
INSERT INTO animations(anime_id, genre_id, studio_id, title, evaluate , release_date, anime_status, rating)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateAnime :exec
UPDATE animations
SET genre_id = ?, studio_id = ?, title = ?, evaluate = ?, release_date = ?, anime_status = ?, rating = ?
WHERE anime_id = ?;

-- name: DeleteAnime :exec
DELETE FROM animations
WHERE anime_id = ?;