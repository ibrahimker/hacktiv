package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ibrahimker/latihan-register/entity"
	"github.com/ibrahimker/latihan-register/service"
)

type RegisterHandlerInterface interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
}

type RegisterHandler struct {
	userService service.UserService
}

func NewRegisterHandler(userService service.UserService) RegisterHandlerInterface {
	return &RegisterHandler{userService: userService}
}

func (h *RegisterHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.registerHandler(w, r)
	}
}

func (h *RegisterHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve request
	decoder := json.NewDecoder(r.Body)
	var req entity.RegisterRequest
	if err := decoder.Decode(&req); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	// throw to service
	resSvc, err := h.userService.Register(r.Context(), &entity.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Age:      req.Age,
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	res := entity.RegisterResponse{
		Id:       resSvc.Id,
		Username: resSvc.Username,
		Email:    resSvc.Email,
		Age:      resSvc.Age,
	}
	jsonData, _ := json.Marshal(&res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
