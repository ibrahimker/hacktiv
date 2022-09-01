package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/hacktiv/latihan/http-server/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	fmt.Println("http server running on localhost:8080")
	// http handler bawaan
	//http.HandleFunc("/register", userSvc.RegisterHandler)
	//http.HandleFunc("/user", userSvc.GetUserHandler)
	//http.ListenAndServe("localhost:8080", nil)

	// with mux
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloWorldHandler)
	r.HandleFunc("/register", userSvc.RegisterHandler)
	r.HandleFunc("/user", userSvc.GetUserHandler)
	r.HandleFunc("/user/{id}", userSvc.GetUserHandler)

	// with gin
	//r := gin.Default()
	//r.GET("/hello", gin.WrapF(helloWorldHandler))
	//r.POST("/register", gin.WrapF(userSvc.RegisterHandler))
	//r.GET("/user", gin.WrapF(userSvc.GetUserHandler))
	//r.GET("/user/:id", gin.WrapF(userSvc.GetUserHandler))

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
