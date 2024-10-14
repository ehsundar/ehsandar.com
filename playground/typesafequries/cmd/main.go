package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ehsundar/ehsandarcom/playground/typesafequries/users"
	"time"
)

type Server struct {
	pool    *pgxpool.Pool
	querier users.Querier
}

func NewServer(
	pool *pgxpool.Pool,
	querier users.Querier,
) Server {
	return Server{
		pool:    pool,
		querier: querier,
	}
}

func (s Server) CreateUser(ctx context.Context, name string, age int32) ([]byte, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = tx.Rollback(queryCtx)
		if err != nil {
			fmt.Println("error rolling back transaction")
		}
	}()

	user, err := s.querier.CreateUser(queryCtx, tx, users.CreateUserParams{
		Name: name,
		Age:  age,
	})

	if err != nil {
		return []byte{}, err
	}

	err = tx.Commit(queryCtx)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("%+v\n", user)), nil
}
