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
	"github.com/ibrahimker/latihan-register/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/api/idtoken"
)

const PORT = ":8080"

func main() {
	postgrespool, perr := newPostgresPool("localhost", "5432", "postgresuser", "postgrespassword", "postgres")
	if perr != nil {
		log.Fatal(perr)
	}

	googleTokenValidator, perr := idtoken.NewValidator(context.Background())
	if perr != nil {
		log.Fatal(perr)
	}

	var cfg config.Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	userHandler := handler.NewUserHandler(postgrespool)
	r.HandleFunc("/users", userHandler.UsersHandler)
	r.HandleFunc("/users/{id}", userHandler.UsersHandler)
	orderHandler := handler.NewOrderHandler(postgrespool)
	r.HandleFunc("/orders", orderHandler.OrderHandler)
	r.HandleFunc("/orders/{id}", orderHandler.OrderHandler)

	authMiddleware := middleware.NewAuthMiddleware(&cfg, googleTokenValidator)
	// r.Use(authMiddleware.AuthBasicMiddleware)
	r.Use(authMiddleware.AuthGoogleIDTokenMiddleware)

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
