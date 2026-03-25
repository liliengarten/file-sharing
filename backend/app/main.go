package main

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://postgres:sharikvadrati@localhost:5432/file-sharing")
	
	if err != nil {
		fmt.Println(err)
	}

	pool.Exec(ctx, `CREATE TABLE USERS(
		id serial PRIMARY KEY,
		first_name varchar(255) not null,
		last_name varchar(255) not null,
		username varchar(255) unique not null,
		password varchar(255) not null
	)`)

	defer pool.Close()
}
