// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: skills.sql

package db

import (
	"context"
)

const addSkill = `-- name: AddSkill :one
INSERT INTO skills (skill)
VALUES
    ($1) RETURNING skill
`

func (q *Queries) AddSkill(ctx context.Context, skill string) (string, error) {
	row := q.db.QueryRow(ctx, addSkill, skill)
	err := row.Scan(&skill)
	return skill, err
}

const deleteSkill = `-- name: DeleteSkill :exec
DELETE FROM
    skills
WHERE
    skill = $1
`

func (q *Queries) DeleteSkill(ctx context.Context, skill string) error {
	_, err := q.db.Exec(ctx, deleteSkill, skill)
	return err
}

const getSkill = `-- name: GetSkill :one
SELECT
    skill
FROM skills
WHERE
    skill = $1
`

func (q *Queries) GetSkill(ctx context.Context, skill string) (string, error) {
	row := q.db.QueryRow(ctx, getSkill, skill)
	err := row.Scan(&skill)
	return skill, err
}

const getSkills = `-- name: GetSkills :many
SELECT
    skill
FROM skills
`

func (q *Queries) GetSkills(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, getSkills)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var skill string
		if err := rows.Scan(&skill); err != nil {
			return nil, err
		}
		items = append(items, skill)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
