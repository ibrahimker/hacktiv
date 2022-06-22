package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ibrahimker/latihan-register/entity"
)

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.loginHandler(w, r)
	}
}

func (h *UserHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
}
