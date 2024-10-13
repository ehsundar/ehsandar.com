package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ehsundar/ehsandarcom/playground/typesafequries/users"
)

type Server struct {
	pool *pgxpool.Pool
}

func NewServer(
	pool *pgxpool.Pool,
) Server {
	return Server{
		pool: pool,
	}
}

func (s Server) CreateUser(ctx context.Context, name string, age int32) ([]byte, error) {
	tx, _ := s.pool.Begin(ctx)
	q := users.New(tx)

	user, err := q.CreateUser(ctx, users.CreateUserParams{
		Name: name,
		Age:  age,
	})

	if err != nil {
		return []byte{}, err
	}

	return []byte(fmt.Sprintf("%+v\n", user)), nil
}
