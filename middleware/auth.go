package middleware

import (
	"net/http"

	"github.com/ibrahimker/latihan-register/config"
)

type AuthMiddlewareIface interface {
	AuthBasicMiddleware(next http.Handler) http.Handler
}

type AuthMiddleware struct {
	cfg *config.Config
}

func NewAuthMiddleware(cfg *config.Config) AuthMiddlewareIface {
	return &AuthMiddleware{cfg: cfg}
}

func (m *AuthMiddleware) AuthBasicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if inputUsername, inputPassword, ok := r.BasicAuth(); !ok || inputUsername != m.cfg.Username || inputPassword != m.cfg.Password {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized request"))
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
