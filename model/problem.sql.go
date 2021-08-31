// Code generated by sqlc. DO NOT EDIT.
// source: problem.sql

package model

import (
	"context"

	"github.com/google/uuid"
)

const deleteProblem = `-- name: DeleteProblem :exec
DELETE FROM problem WHERE id = $1
`

func (q *Queries) DeleteProblem(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteProblemStmt, deleteProblem, id)
	return err
}

const getProblemByID = `-- name: GetProblemByID :one
SELECT id, name, description, category, difficulty, grader_code, create_at, update_at FROM problem WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProblemByID(ctx context.Context, id uuid.UUID) (Problem, error) {
	row := q.queryRow(ctx, q.getProblemByIDStmt, getProblemByID, id)
	var i Problem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.Difficulty,
		&i.GraderCode,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getProblems = `-- name: GetProblems :many
SELECT id, name, description, category, difficulty, grader_code, create_at, update_at FROM problem
`

func (q *Queries) GetProblems(ctx context.Context) ([]Problem, error) {
	rows, err := q.query(ctx, q.getProblemsStmt, getProblems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Problem
	for rows.Next() {
		var i Problem
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.Difficulty,
			&i.GraderCode,
			&i.CreateAt,
			&i.UpdateAt,
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

const insertProblem = `-- name: InsertProblem :one
INSERT INTO problem (name, description, category, difficulty, grader_code) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, description, category, difficulty, grader_code, create_at, update_at
`

type InsertProblemParams struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Difficulty  DifficultyType `json:"difficulty"`
	GraderCode  string         `json:"grader_code"`
}

func (q *Queries) InsertProblem(ctx context.Context, arg InsertProblemParams) (Problem, error) {
	row := q.queryRow(ctx, q.insertProblemStmt, insertProblem,
		arg.Name,
		arg.Description,
		arg.Category,
		arg.Difficulty,
		arg.GraderCode,
	)
	var i Problem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.Difficulty,
		&i.GraderCode,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const updateProblem = `-- name: UpdateProblem :one
UPDATE problem SET name = $2, description = $3, category = $4, difficulty = $5, grader_code = $6 WHERE id = $1 RETURNING id, name, description, category, difficulty, grader_code, create_at, update_at
`

type UpdateProblemParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Difficulty  DifficultyType `json:"difficulty"`
	GraderCode  string         `json:"grader_code"`
}

func (q *Queries) UpdateProblem(ctx context.Context, arg UpdateProblemParams) (Problem, error) {
	row := q.queryRow(ctx, q.updateProblemStmt, updateProblem,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Category,
		arg.Difficulty,
		arg.GraderCode,
	)
	var i Problem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.Difficulty,
		&i.GraderCode,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
