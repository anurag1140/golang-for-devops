package middleware

import (
	"net/http"

	"golang-for-devops/internal/auth"
)

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims, ok := r.Context().
				Value(auth.UserContextKey).(*auth.Claims)

			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			authorized := false

			for _, role := range allowedRoles {

				if claims.Role == role {
					authorized = true
					break
				}
			}

			if !authorized {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
