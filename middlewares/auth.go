package middlewares

import (
	"blog-platform/utils"
	"net/http"
	"strings"

	"github.com/gorilla/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		tokenStr := strings.Split(authHeader, " ")[1]
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		context.Set(r, "username", claims.Username)
		context.Set(r, "role", claims.Role)
		next.ServeHTTP(w, r)
	})
}

func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := context.Get(r, "role").(string)
		if role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
