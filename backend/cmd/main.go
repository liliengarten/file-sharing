package main

import (
	"context"
	"database/sql"
	"liliengarten/filesharing/internal/middlewares"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"liliengarten/filesharing/internal/handlers"
	"liliengarten/filesharing/internal/repository"
	"liliengarten/filesharing/internal/service"
)

func setupRoutes(mux *http.ServeMux, userHandler *handlers.UserHandler, pinHandler *handlers.PinHandler) {
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	mux.HandleFunc("GET /pins", middlewares.AuthMiddleware(pinHandler.Index))
	mux.HandleFunc("POST /pins", middlewares.AuthMiddleware(pinHandler.Add))
	mux.HandleFunc("PATCH /pins/{id}", middlewares.AuthMiddleware(pinHandler.Update))
	mux.HandleFunc("DELETE /pins/{id}", middlewares.AuthMiddleware(pinHandler.Remove))

	/*TODO:
	лайк
	подписка на пользователя

	создание доски
	добавление пина на доску
	удаление пина с доски

	добавление пользователя на доску
	удаление пользователя с доски


	функционал сохраненных пинов не нужен, лайков достаточно
	*/
}

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

	userRepo := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	pinRepo := repository.NewPinRepository(pool)
	pinService := service.NewPinService(pinRepo)
	pinHandler := handlers.NewPinHandler(pinService)

	mux := http.NewServeMux()
	setupRoutes(mux, userHandler, pinHandler)
	http.ListenAndServe(":8080", mux)
}
