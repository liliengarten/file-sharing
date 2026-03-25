package main

import (
	"log"
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib" 
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:sharikvadrati@localhost:5432/file-sharing")
	if err != nil {
		log.Fatal(err)
	}

	goose.SetDialect("postgres")

	err = goose.Up(db, "./../migrations")
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://postgres:sharikvadrati@localhost:5432/file-sharing")
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()
}
