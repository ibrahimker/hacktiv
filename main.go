package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/latihan-register/config"
	"github.com/ibrahimker/latihan-register/handler"
	"github.com/ibrahimker/latihan-register/repository"
	"github.com/ibrahimker/latihan-register/service"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v4/pgxpool"
)

const PORT = ":8080"

func main() {
	var cfg config.Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	postgrespool, perr := newPostgresPool("localhost", "5432", cfg.Username, cfg.Password, "postgres")
	if perr != nil {
		log.Fatal(perr)
	}

	r := mux.NewRouter()

	// init repository
	userRepo := repository.NewUserRepo(postgrespool)

	// init service
	userService := service.NewUserSvc(userRepo)

	// init handler
	loginHandler := handler.NewLoginHandler(userService)
	registerHandler := handler.NewRegisterHandler(userService)
	// userHandler := handler.NewUserHandler(postgrespool)

	// setup route
	// r.HandleFunc("/users", userHandler.UsersHandler)
	// r.HandleFunc("/users/{id}", userHandler.UsersHandler)
	r.HandleFunc("/users/login", loginHandler.LoginHandler)
	r.HandleFunc("/users/register", registerHandler.RegisterHandler)
	// orderHandler := handler.NewOrderHandler(postgrespool)
	// r.HandleFunc("/orders", orderHandler.OrderHandler)
	// r.HandleFunc("/orders/{id}", orderHandler.OrderHandler)

	// authMiddleware := middleware.NewAuthMiddleware(&cfg)
	// r.Use(authMiddleware.AuthBasicMiddleware)
	// r.Use(authMiddleware.AuthGoogleIDTokenMiddleware)

	const htmlPath = "static/web.html"
	const jsonPath = "static/weather.json"

	go handler.GenerateToJson()

	r.HandleFunc("/assignment3", handler.Assignment3Handler)

	fmt.Println("Now listening on port 0.0.0.0" + PORT)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// newPostgresPool builds a pool of pgx client.
func newPostgresPool(host, port, user, password, name string) (*pgxpool.Pool, error) {
	connCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
	return pgxpool.Connect(context.Background(), connCfg)
}
