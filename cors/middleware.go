package cors

import (
	"fmt"
	"net/http"
)

// Middleware :
func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		fmt.Println("before executing controller")

		handler.ServeHTTP(w, r)

		fmt.Println("after executing controller")
	})
}
