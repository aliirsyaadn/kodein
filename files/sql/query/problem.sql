-- name: GetProblems :many
SELECT * FROM problem;

-- name: GetProblemByID :one
SELECT * FROM problem WHERE id = $1 LIMIT 1;

-- name: InsertProblem :one
INSERT INTO problem (name, description, category, difficulty, grader_code) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteProblem :exec
DELETE FROM problem WHERE id = $1;

-- name: UpdateProblem :one
UPDATE problem SET name = $2, description = $3, category = $4, difficulty = $5, grader_code = $6 WHERE id = $1 RETURNING *;