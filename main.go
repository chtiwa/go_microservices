package main

import (
	"fmt"
	"net/http"
)

func main() {
	// saving it as a pointer and taking the memory address
	server := &http.Server{
		Addr:    ":8000",
		Handler: http.HandlerFunc(basicHandler),
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
