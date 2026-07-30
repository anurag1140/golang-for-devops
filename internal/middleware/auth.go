package middleware

import (
	"context"
	"net/http"
	"strings"

	"golang-for-devops/internal/auth"

	apierrors "golang-for-devops/internal/errors"
)

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			apierrors.WriteError(
				w,
				apierrors.Unauthorized(apierrors.CodeUnauthorized),
			)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			apierrors.WriteError(
				w,
				apierrors.Unauthorized(apierrors.CodeUnauthorized),
			)
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		claims, err := auth.ValidateToken(tokenString)

		if err != nil {
			apierrors.WriteError(
				w,
				err,
			)
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
