-- name: GetUser :one
SELECT * FROM users
WHERE user_id = ?;

-- name: AddUser :exec
INSERT INTO users(user_id, user_name, email, user_pwd, avatar_url)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users
SET user_name = ?, email = ?, user_pwd = ?, avatar_url = ?
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;