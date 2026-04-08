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

func setupRoutes(mux *http.ServeMux, userHandler *handlers.UserHandler, pinHandler *handlers.PinHandler, boardHandler *handlers.BoardHandler) {
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	mux.HandleFunc("GET /profile/{id}", userHandler.GetProfile)

	mux.HandleFunc("GET /pins", middlewares.AuthMiddleware(pinHandler.Index))
	mux.HandleFunc("GET /pins/{id}", middlewares.AuthMiddleware(pinHandler.GetPin))
	mux.HandleFunc("POST /pins", middlewares.AuthMiddleware(pinHandler.Add))
	mux.HandleFunc("PATCH /pins/{id}", middlewares.AuthMiddleware(pinHandler.Update))
	mux.HandleFunc("DELETE /pins/{id}", middlewares.AuthMiddleware(pinHandler.Remove))

	mux.HandleFunc("POST /pins/{id}/like", middlewares.AuthMiddleware(pinHandler.LikePin))
	mux.HandleFunc("DELETE /pins/{id}/like", middlewares.AuthMiddleware(pinHandler.UnlikePin))

	mux.HandleFunc("GET /boards", middlewares.AuthMiddleware(boardHandler.Index))
	mux.HandleFunc("GET /boards/{id}", middlewares.AuthMiddleware(boardHandler.GetBoard))
	mux.HandleFunc("POST /boards", middlewares.AuthMiddleware(boardHandler.Create))
	mux.HandleFunc("DELETE /boards/{id}", middlewares.AuthMiddleware(boardHandler.Remove))

	mux.HandleFunc("GET /boards/{id}/pins", middlewares.AuthMiddleware(boardHandler.GetPins))
	mux.HandleFunc("POST /boards/{board_id}/pins/{pin_id}", middlewares.AuthMiddleware(boardHandler.AddPin))
	mux.HandleFunc("DELETE /boards/{board_id}/pins/{pin_id}", middlewares.AuthMiddleware(boardHandler.RemovePin))

	mux.HandleFunc("GET /boards/{id}/authors", middlewares.AuthMiddleware(boardHandler.GetAuthors))
	mux.HandleFunc("POST /boards/{board_id}/authors/{user_id}", middlewares.AuthMiddleware(boardHandler.AddAuthor))
	mux.HandleFunc("DELETE /boards/{board_id}/authors/{user_id}", middlewares.AuthMiddleware(boardHandler.RemoveAuthor))
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

	boardRepo := repository.NewBoardRepository(pool)
	boardService := service.NewBoardService(boardRepo, userRepo)
	boardHandler := handlers.NewBoardHandler(boardService)

	mux := http.NewServeMux()
	setupRoutes(mux, userHandler, pinHandler, boardHandler)
	http.ListenAndServe(":8080", mux)
}
