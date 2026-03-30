package main

import (
	"net/http"
	"log"
	"context"
	"database/sql"
	
	_ "github.com/jackc/pgx/v5/stdlib" 
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"

	"liliengarten/filesharing/internal/handlers"
	"liliengarten/filesharing/internal/service"
	"liliengarten/filesharing/internal/repository"
)

func setupRoutes(userHandler *handlers.UserHandler) {
	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)
}

func main() {
	//Миграции
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

	//Подклбчение к БД
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:sharikvadrati@localhost:5432/file-sharing")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	
	//Инициализация репозиториев, сервисов и хендлеров
	repo := repository.NewUserRepository(pool)
	service := service.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	//Запуск роутов и сервера
	setupRoutes(handler)
	http.ListenAndServe(":8080", nil)
}
