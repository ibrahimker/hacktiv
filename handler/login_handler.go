package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ibrahimker/latihan-register/entity"
	"github.com/ibrahimker/latihan-register/service"
)

type LoginHandlerInterface interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
}

type LoginHandler struct {
	userService service.UserService
}

func NewLoginHandler(userService service.UserService) LoginHandlerInterface {
	return &LoginHandler{userService: userService}
}

func (h *LoginHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.loginHandler(w, r)
	}
}

func (l *LoginHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
}
