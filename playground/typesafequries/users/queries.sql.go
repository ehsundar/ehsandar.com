// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package users

import (
	"context"
)

const createUser = `-- name: CreateUser :one
insert into users (name, age)
values ($1, $2)
returning id, name, age
`

func (q *Queries) CreateUser(ctx context.Context, db DBTX, name string, age int32) (User, error) {
	row := db.QueryRow(ctx, createUser, name, age)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Age)
	return i, err
}
