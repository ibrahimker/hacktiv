package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/ibrahimker/latihan-register/config"
	"google.golang.org/api/idtoken"
)

type AuthMiddlewareIface interface {
	AuthBasicMiddleware(next http.Handler) http.Handler
	AuthGoogleIDTokenMiddleware(next http.Handler) http.Handler
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

func (m *AuthMiddleware) AuthGoogleIDTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authHeader := strings.Split(r.Header.Get("Authorization"), " "); len(authHeader) == 2 && authHeader[0] == "Bearer" {
			idToken := authHeader[1]
			idTokenPayload, err := idtoken.Validate(r.Context(), idToken, m.cfg.GoogleClientID)
			if err != nil {
				log.Println("Error when validate id token", err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized request"))
				return
			}
			log.Println(idTokenPayload)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized request"))
			return
		}
	})
}
