package middleware

import (
	"context"
	"net/http"
	"strings"

	"golang-for-devops/internal/auth"
)

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		claims, err := auth.ValidateToken(tokenString)

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			auth.UserContextKey,
			claims,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
