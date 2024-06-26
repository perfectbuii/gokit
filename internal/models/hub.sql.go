// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: hub.sql

package models

import (
	"context"
	"database/sql"
)

const createHub = `-- name: CreateHub :one
INSERT INTO hubs(id,name,location_id)
values($1,$2,$3)
Returning id, name, location_id, created_at, updated_at, deleted_at
`

type CreateHubParams struct {
	ID         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	LocationID string `db:"location_id" json:"location_id"`
}

func (q *Queries) CreateHub(ctx context.Context, arg CreateHubParams) (*Hub, error) {
	row := q.db.QueryRowContext(ctx, createHub, arg.ID, arg.Name, arg.LocationID)
	var i Hub
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteHub = `-- name: DeleteHub :one
DELETE from hubs where id = $1
Returning id, name, location_id, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteHub(ctx context.Context, id string) (*Hub, error) {
	row := q.db.QueryRowContext(ctx, deleteHub, id)
	var i Hub
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const findHubByID = `-- name: FindHubByID :one
SELECT id, name, location_id, created_at, updated_at, deleted_at FROM hubs
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindHubByID(ctx context.Context, id string) (*Hub, error) {
	row := q.db.QueryRowContext(ctx, findHubByID, id)
	var i Hub
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getListHub = `-- name: GetListHub :many
SELECT id, name, location_id, created_at, updated_at, deleted_at FROM hubs
offset $1 limit $2
`

type GetListHubParams struct {
	Offset int32 `db:"offset" json:"offset"`
	Limit  int32 `db:"limit" json:"limit"`
}

func (q *Queries) GetListHub(ctx context.Context, arg GetListHubParams) ([]*Hub, error) {
	rows, err := q.db.QueryContext(ctx, getListHub, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Hub
	for rows.Next() {
		var i Hub
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.LocationID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchHub = `-- name: SearchHub :many
SELECT id, name, location_id, created_at, updated_at, deleted_at from hubs
where name like ('%' || $1 || '%')
`

func (q *Queries) SearchHub(ctx context.Context, dollar_1 sql.NullString) ([]*Hub, error) {
	rows, err := q.db.QueryContext(ctx, searchHub, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Hub
	for rows.Next() {
		var i Hub
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.LocationID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHub = `-- name: UpdateHub :one
UPDATE hubs set name = $1
where id = $2
Returning id, name, location_id, created_at, updated_at, deleted_at
`

type UpdateHubParams struct {
	Name string `db:"name" json:"name"`
	ID   string `db:"id" json:"id"`
}

func (q *Queries) UpdateHub(ctx context.Context, arg UpdateHubParams) (*Hub, error) {
	row := q.db.QueryRowContext(ctx, updateHub, arg.Name, arg.ID)
	var i Hub
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
