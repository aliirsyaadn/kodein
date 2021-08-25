-- name: GetMembers :many
SELECT * FROM member;

-- name: GetMemberByID :one
SELECT * FROM member WHERE id = $1 LIMIT 1;

-- name: GetMemberByUsername :one
SELECT * FROM member WHERE username = $1 LIMIT 1;

-- name: InsertMember :one
INSERT INTO member (name, username, password, email) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteMember :exec
DELETE FROM member WHERE id = $1;

-- name: UpdateMember :one
UPDATE member SET name = $2, username = $3, email = $4 WHERE id = $1 RETURNING *;

-- name: UpdateSosmedMember :one
UPDATE member SET twitter = $2, github = $3, linkedin = $4 WHERE id = $1 RETURNING *;

-- name: UpdatePasswordMember :one
UPDATE member SET password = $2 WHERE id = $1 RETURNING *;