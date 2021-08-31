-- name: GetAttemptsByMemberID :many
SELECT * FROM attempt WHERE member_id = $1;

-- name: GetAttemptByID :one
SELECT * FROM attempt WHERE id = $1 LIMIT 1;

-- name: InsertAttempt :one
INSERT INTO attempt (member_id, problem_id, language, is_solved, score, code) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: DeleteAttempt :exec
DELETE FROM attempt WHERE id = $1;

-- name: UpdateAttempt :one
UPDATE attempt SET language = $2, is_solved = $3, score = $4, code = $5 WHERE id = $1 RETURNING *;