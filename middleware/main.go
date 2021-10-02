package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/dashboard", checkToken(dashboard))
	http.ListenAndServe(":3000", nil)
}

func checkToken(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(response, "No Token Provided", 401)
			return
		}

		next.ServeHTTP(response, request)
	}
  }

func dashboard(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))

	if err != nil {
		http.Error(response, "Internal Server Error", 500)
		return
	}
	
	fmt.Fprintf(response, "ID %d...", id)
}