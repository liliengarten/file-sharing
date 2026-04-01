package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"liliengarten/filesharing/internal/handlers"
	"liliengarten/filesharing/internal/repository"
	"liliengarten/filesharing/internal/service"
)

func setupRoutes(userHandler *handlers.UserHandler, pinHandler *handlers.PinHandler) {
	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	http.HandleFunc("/pins", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pinHandler.Index(w, r)
		case http.MethodPost:
			pinHandler.Add(w, r)
		case http.MethodPatch:
			pinHandler.Update(w, r)
		case http.MethodDelete:
			pinHandler.Remove(w, r)
		}
	})

	http.H
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
	userRepo := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	pinRepo := repository.NewPinRepository(pool)
	pinService := service.NewPinService(pinRepo)
	pinHandler := handlers.NewPinHandler(pinService)

	//Запуск роутов и сервера
	setupRoutes(userHandler, pinHandler)
	http.ListenAndServe(":8080", nil)
}
