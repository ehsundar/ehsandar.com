-- name: CreateUser :one
insert into users (name, age)
values ($1, $2)
returning id, name, age;
