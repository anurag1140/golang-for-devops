package middleware

import (
	"fmt"
	"net/http"
)

func Logging(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Request Received")

		fmt.Println("Method:", r.Method)

		fmt.Println("Path:", r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Println("Request Completed")
	})
}
