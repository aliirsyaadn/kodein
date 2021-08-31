-- name: GetProjectsByMemberID :many
SELECT * FROM project WHERE member_id = $1;

-- name: GetProjectByID :one
SELECT * FROM project WHERE id = $1 LIMIT 1;

-- name: InsertProject :one
INSERT INTO project (member_id, name, description, technology, url) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteProject :exec
DELETE FROM project WHERE id = $1;

-- name: UpdateProject :one
UPDATE project SET name = $2, description = $3, technology = $4, url = $5 WHERE id = $1 RETURNING *;