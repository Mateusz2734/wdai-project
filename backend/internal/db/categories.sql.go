// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: categories.sql

package db

import (
	"context"
)

const addCategory = `-- name: AddCategory :one
INSERT INTO categories (category)
VALUES
    ($1) RETURNING category
`

func (q *Queries) AddCategory(ctx context.Context, category string) (string, error) {
	row := q.db.QueryRow(ctx, addCategory, category)
	err := row.Scan(&category)
	return category, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM
    categories
WHERE
    category = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, category string) error {
	_, err := q.db.Exec(ctx, deleteCategory, category)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT
    category
FROM categories
`

func (q *Queries) GetCategories(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, getCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		items = append(items, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoriesBySkill = `-- name: GetCategoriesBySkill :many
SELECT
    category
FROM skill_categories
WHERE
    skill = $1
`

func (q *Queries) GetCategoriesBySkill(ctx context.Context, skill string) ([]string, error) {
	rows, err := q.db.Query(ctx, getCategoriesBySkill, skill)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		items = append(items, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
