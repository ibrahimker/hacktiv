package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/latihan-register/handler"
	"github.com/jackc/pgx/v4/pgxpool"
)

const PORT = ":8080"

func main() {
	postgrespool, perr := newPostgresPool("localhost", "5432", "postgresuser", "postgrespassword", "postgres")
	if perr != nil {
		log.Fatal(perr)
	}

	r := mux.NewRouter()
	userHandler := handler.NewUserHandler(postgrespool)
	r.HandleFunc("/users", userHandler.UsersHandler)
	r.HandleFunc("/users/{id}", userHandler.UsersHandler)

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
