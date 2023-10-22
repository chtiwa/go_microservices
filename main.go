package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", basicHandler)
	// saving it as a pointer and taking the memory address
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}

}

// r is a pointer of http.Request
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
