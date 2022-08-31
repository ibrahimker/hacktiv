package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Nama string `json:"nama"`
}

type userService struct {
	db []*User
}

type UserIface interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	Register(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) UserIface {
	return &userService{
		db: db,
	}
}

func (u *userService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + "berhasil didaftarkan"
}

func (u *userService) GetUser() []*User {
	return u.db
}

func (u *userService) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			fmt.Println("error data user")
			return
		}
		u.Register(&user)
		w.Write([]byte(user.Nama + " berhasil didaftarkan"))
	} else {
		w.Write([]byte("invalid http method"))
	}
}

func (u *userService) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users := u.GetUser()
	data, _ := json.Marshal(users)
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}
