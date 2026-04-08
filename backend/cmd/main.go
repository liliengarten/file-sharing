package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	"liliengarten/filesharing/internal/handlers"
	"liliengarten/filesharing/internal/middlewares"
	"liliengarten/filesharing/internal/repository"
	"liliengarten/filesharing/internal/service"
)

func setupRoutes(mux *http.ServeMux, userHandler *handlers.UserHandler, pinHandler *handlers.PinHandler, boardHandler *handlers.BoardHandler) {
	protectedMux := http.NewServeMux()

	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)
	protectedMux.HandleFunc("GET /profile/{id}", userHandler.GetProfile)

	protectedMux.HandleFunc("GET /pins", pinHandler.Index)
	protectedMux.HandleFunc("GET /pins/{id}", pinHandler.GetPin)
	protectedMux.HandleFunc("POST /pins", pinHandler.Add)
	protectedMux.HandleFunc("PATCH /pins/{id}", pinHandler.Update)
	protectedMux.HandleFunc("DELETE /pins/{id}", pinHandler.Remove)

	protectedMux.HandleFunc("GET /likes", userHandler.GetLikes)
	protectedMux.HandleFunc("POST /pins/{id}/like", pinHandler.LikePin)
	protectedMux.HandleFunc("DELETE /pins/{id}/like", pinHandler.UnlikePin)

	protectedMux.HandleFunc("GET /boards", boardHandler.Index)
	protectedMux.HandleFunc("GET /boards/{id}", boardHandler.GetBoard)
	protectedMux.HandleFunc("POST /boards", boardHandler.Create)
	protectedMux.HandleFunc("DELETE /boards/{id}", boardHandler.Remove)

	protectedMux.HandleFunc("GET /boards/{id}/pins", boardHandler.GetPins)
	protectedMux.HandleFunc("POST /boards/{board_id}/pins/{pin_id}", boardHandler.AddPin)
	protectedMux.HandleFunc("DELETE /boards/{board_id}/pins/{pin_id}", boardHandler.RemovePin)

	protectedMux.HandleFunc("GET /boards/{id}/authors", boardHandler.GetAuthors)
	protectedMux.HandleFunc("POST /boards/{board_id}/authors/{user_id}", boardHandler.AddAuthor)
	protectedMux.HandleFunc("DELETE /boards/{board_id}/authors/{user_id}", boardHandler.RemoveAuthor)

	mux.Handle("/", middlewares.AuthMiddleware(protectedMux))
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatal(err)
	}

	err = goose.SetDialect("postgres")
	if err != nil {
		log.Fatal(err)
	}
	err = goose.Up(db, "./../migrations")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	pinRepo := repository.NewPinRepository(pool)
	pinService := service.NewPinService(pinRepo)
	pinHandler := handlers.NewPinHandler(pinService)

	userRepo := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepo, pinRepo)
	userHandler := handlers.NewUserHandler(userService)

	boardRepo := repository.NewBoardRepository(pool)
	boardService := service.NewBoardService(boardRepo, userRepo)
	boardHandler := handlers.NewBoardHandler(boardService)

	mux := http.NewServeMux()
	setupRoutes(mux, userHandler, pinHandler, boardHandler)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
